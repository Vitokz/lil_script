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
	"github.com/tendermint/tendermint/libs/sync"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"time"

	tmjson "github.com/tendermint/tendermint/libs/json"
	tmtypes "github.com/tendermint/tendermint/types"
	//coretypes "github.com/tendermint/tendermint/rpc/core/types"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
	//"time"
)

type event string

const (
	TxEvent event = tmtypes.EventTx
	//newTxSub event = "new_tx_subscriber"
)

type Worker struct {
	ctx    context.Context
	logger log.Logger
	cdc    *params.EncodingConfig
	Rpc    *rpc.Client
	Web3   *web3.Client
	mm     types.ModuleManager
	db     *db.DB

	topicChans map[string]chan coretypes.ResultEvent
	mux        *sync.RWMutex

	eventsSubs []event
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

	topicChans := make(map[string]chan coretypes.ResultEvent)

	return &Worker{
		ctx:    context.Background(),
		cdc:    cdc,
		logger: logger,
		Web3:   web3Client,
		Rpc:    rpcCLient,
		mm:     mm,
		db:     db,

		topicChans: topicChans,
		mux:        new(sync.RWMutex),

		eventsSubs: make([]event, 0),
	}, nil
}

func (w *Worker) StartWorker(events ...event) {
	//re, error := w.Rpc.TmRpc.Subscribe(w.ctx, newTxSub, tmtypes.EventTx)
	for _, v := range events {
		event := string(v)
		err := w.Rpc.TmWsClient.Subscribe(w.ctx, event)
		if err != nil {
			panic(err.Error())
		}

		ch := make(chan coretypes.ResultEvent)
		w.topicChans[event] = ch
	}

	go w.startReadNewTxEvents()

	w.eventsSubs = events
}

func (w *Worker) startReadNewTxEvents() {
	ch, ok := w.topicChans[string(TxEvent)]
	if !ok {
		return
	}

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
					continue //panic(err) //TODO: check error
				}
				err = handler(context.WithValue(context.Background(), "height", dataTx.Height), w.Web3, w.db, w.logger, msg)
				if err != nil {
					ch <- event // do retry //TODO: нужно ли retry
				}
			}
		default:
			continue
		}
	}
}

func (w *Worker) consumeEvents() {
	for {
		for rpcResp := range w.Rpc.TmWsClient.ResponsesCh {
			var ev coretypes.ResultEvent

			if rpcResp.Error != nil {
				time.Sleep(5 * time.Second)
				continue
			} else if err := tmjson.Unmarshal(rpcResp.Result, &ev); err != nil {
				w.logger.Error("failed to JSON unmarshal ResponsesCh result event", "error", err.Error())
				continue
			}

			if len(ev.Query) == 0 {
				// skip empty responses
				continue
			}

			w.mux.RLock()
			ch, ok := w.topicChans[ev.Query]
			w.mux.RUnlock()
			if !ok {
				w.logger.Debug("channel for subscription not found", "topic", ev.Query)
				w.logger.Debug("list of available channels", "channels", w.eventsSubs)
				continue
			}

			// gracefully handle lagging subscribers
			t := time.NewTimer(time.Second)
			select {
			case <-t.C:
				w.logger.Debug("dropped event during lagging subscription", "topic", ev.Query)
			case ch <- ev:
			}
		}

		time.Sleep(time.Second)
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
