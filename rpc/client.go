package rpc

import (
	"github.com/Vitokz/lil_script/config"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/tendermint/tendermint/libs/log"
	tmRpc "github.com/tendermint/tendermint/rpc/client/http"
	tmJsonRpcClient "github.com/tendermint/tendermint/rpc/jsonrpc/client"
	"net/http"
	"time"
)

type Client struct {
	cdc        *params.EncodingConfig
	logger     *log.Logger
	TmRpc      *tmRpc.HTTP
	TmWsClient *tmJsonRpcClient.WSClient
}

func NewClient(cfg config.ConfigI, cdc *params.EncodingConfig, logger log.Logger) (*Client, error) {
	httpClient := &http.Client{}
	tmRpcClient, err := tmRpc.NewWithClient(cfg.GetTendermintRPC(), cfg.GetTendermintRPC(), httpClient)
	if err != nil {
		return nil, err
	}

	endpoint := "/websocket"
	tmWsClient, err := tmJsonRpcClient.NewWS(cfg.GetTendermintRPC(), endpoint,
		tmJsonRpcClient.MaxReconnectAttempts(256),
		tmJsonRpcClient.ReadWait(120*time.Second),
		tmJsonRpcClient.WriteWait(120*time.Second),
		tmJsonRpcClient.PingPeriod(50*time.Second),
		tmJsonRpcClient.OnReconnect(func() {
			logger.Debug("Client RPC reconnects to Tendermint WS", "address", cfg.GetTendermintRPC()+endpoint)
		}))
	if err != nil {
		return nil, err
	}

	return &Client{
		TmRpc:      tmRpcClient,
		TmWsClient: tmWsClient,
		cdc:        cdc,
	}, nil
}
