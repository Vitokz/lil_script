package worker

import (
	"context"
	"fmt"
	"github.com/Vitokz/lil_script/config"
	"github.com/Vitokz/lil_script/db"
	"github.com/Vitokz/lil_script/rpc"
	"github.com/Vitokz/lil_script/worker/types"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	web3 "github.com/ethereum/go-ethereum/ethclient"
	"github.com/tendermint/tendermint/libs/log"
	tmSync "github.com/tendermint/tendermint/libs/sync"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"sync"

	tmtypes "github.com/tendermint/tendermint/types"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
)

var (
	TxEvent = tmtypes.QueryForEvent(tmtypes.EventTx).String()
)

type Worker struct {
	ctx    context.Context
	logger log.Logger
	cdc    *params.EncodingConfig
	Rpc    *rpc.Client
	Web3   *web3.Client
	mm     types.ModuleManager
	db     *db.DB
	wg     sync.WaitGroup

	topicChans map[string]<-chan coretypes.ResultEvent
	mux        *tmSync.RWMutex

	eventsSubs []string
}

func NewWorker(cfg config.ConfigI, cdc *params.EncodingConfig, logger log.Logger) (*Worker, error) {
	web3Client, err := web3.Dial(cfg.GetEthereumJsonRPC())
	if err != nil {
		return nil, err
	}

	rpcCLient, err := rpc.NewClient(cfg, cdc, logger)
	if err != nil {
		return nil, err
	}

	db, err := db.NewDB(cfg.GetScoutPgDsn())
	if err != nil {
		return nil, err
	}

	mm := types.NewModuleManager()

	topicChans := make(map[string]<-chan coretypes.ResultEvent)

	return &Worker{
		ctx:    context.Background(),
		cdc:    cdc,
		logger: logger,
		Web3:   web3Client,
		Rpc:    rpcCLient,
		mm:     mm,
		db:     db,
		wg:     sync.WaitGroup{},

		topicChans: topicChans,
		mux:        new(tmSync.RWMutex),

		eventsSubs: make([]string, 0),
	}, nil
}

func (w *Worker) StartWorker(events ...string) {
	for _, event := range events {
		re, err := w.Rpc.TmRpc.Subscribe(w.ctx, event, event)
		if err != nil {
			panic(err.Error())
		}

		w.logger.Info(fmt.Sprintf("Subscribe to event: %s", event))

		w.topicChans[event] = re
	}

	w.wg.Add(1)
	go w.startReadNewTxEvents()

	w.eventsSubs = events
	w.wg.Wait()
}

func (w *Worker) startReadNewTxEvents() {
	ch, ok := w.topicChans[TxEvent]
	if !ok {
		return
	}
	defer w.wg.Done()

	w.logger.Info(fmt.Sprintf("start worker for event: %s", TxEvent))
	for {
		select {
		case event, ok := <-ch:
			if !ok {
				return
			}

			dataTx, ok := event.Data.(tmtypes.EventDataTx)
			if !ok {
				w.logger.Debug("event data type mismatch", "type", fmt.Sprintf("%T", event.Data))
				continue
			}

			w.logger.Info(fmt.Sprintf("consume event with type %s and height %d", event.Query, dataTx.Height))

			tx, err := w.cdc.TxConfig.TxDecoder()(dataTx.Tx)
			if err != nil {
				panic(err.Error())
			}

			ethTxs, err := w.tmTxToEthTx(tx)
			if err == nil && len(ethTxs) == 0 {
				// is ethereum tx
				continue
			}

			for _, msg := range tx.GetMsgs() {
				handler, err := w.mm.GetMsgHandler(sdkTypes.MsgTypeURL(msg))
				if err != nil {
					continue //panic(err) //TODO: handle error
				}
				err = handler(context.WithValue(context.Background(), "height", dataTx.Height), w.Web3, w.db, w.logger, msg)
				if err != nil {
					w.logger.Error(fmt.Sprintf("failed to handle transaction")) //TODO: make retry
				}
			}
		default:
			continue
		}
	}
}

func (w *Worker) tmTxToEthTx(tx sdkTypes.Tx) ([]*evmtypes.MsgEthereumTx, error) {
	ethTxs := make([]*evmtypes.MsgEthereumTx, len(tx.GetMsgs()))
	for i, msg := range tx.GetMsgs() {
		ethTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return nil, fmt.Errorf("invalid message type %T, expected %T", msg, &evmtypes.MsgEthereumTx{})
		}
		ethTxs[i] = ethTx
	}
	return ethTxs, nil
}
