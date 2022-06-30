package rpc

import (
	"github.com/Vitokz/lil_script/config"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/tendermint/tendermint/libs/log"
	//tmRpc "github.com/tendermint/tendermint/rpc/client/http"
	jsonrpcclient "github.com/tendermint/tendermint/rpc/jsonrpc/client"
	"net/http"
	"time"
)

type Client struct {
	cdc    *params.EncodingConfig
	logger *log.Logger
	//TmRpc  *tmRpc.HTTP
	TmJRpc     *jsonrpcclient.Client
	TmWsClient *jsonrpcclient.WSClient
}

func NewClient(cfg config.ConfigI, cdc *params.EncodingConfig, logger log.Logger) (*Client, error) {
	httpClient := &http.Client{}
	//tmRpcClient, err := tmRpc.NewWithClient(cfg.GetTendermintRPC(), "/websocket", httpClient)
	//if err != nil {
	//	return nil, err
	//}
	//if err != nil {
	//	logger.Error(
	//		"Tendermint client could not start",
	//		"address", cfg.GetTendermintRPC(),
	//		"error", err,
	//	)
	//	return nil, err
	//}

	logger.Info("start connect to tm json rpc")
	jrpcCLient, err := jsonrpcclient.NewWithHTTPClient(cfg.GetTendermintRPC(), httpClient)
	if err != nil {
		logger.Error(
			"failed to connecttendermint json rpc",
			"address", cfg.GetTendermintRPC(),
			"error", err,
		)
		return nil, err
	}
	logger.Info("tm json rpc connected!")

	//logger.Info("tm client started!")

	logger.Info("start connect to tm websocket")
	endpoint := "/websocket"
	tmWsClient, err := jsonrpcclient.NewWS(cfg.GetTendermintWebsocket(), endpoint,
		jsonrpcclient.MaxReconnectAttempts(256),
		jsonrpcclient.ReadWait(120*time.Second),
		jsonrpcclient.WriteWait(120*time.Second),
		jsonrpcclient.PingPeriod(50*time.Second),
		jsonrpcclient.OnReconnect(func() {
			logger.Debug("Client RPC reconnects to Tendermint WS", "address", cfg.GetTendermintWebsocket()+endpoint)
		}))
	if err != nil {
		return nil, err
	} else if err := tmWsClient.OnStart(); err != nil {
		logger.Error(
			"Tendermint WS client could not start",
			"address", cfg.GetTendermintWebsocket()+endpoint,
			"error", err,
		)
		return nil, err
	}
	logger.Info("websocket client connected!")

	return &Client{
		//TmRpc: jrpcCLient,
		TmJRpc:     jrpcCLient,
		TmWsClient: tmWsClient,
		cdc:        cdc,
	}, nil
}
