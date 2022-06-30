package rpc

import (
	"github.com/Vitokz/lil_script/config"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/tendermint/tendermint/libs/log"
	tmRpc "github.com/tendermint/tendermint/rpc/client/http"
	"net/http"
)

type Client struct {
	cdc    *params.EncodingConfig
	logger *log.Logger
	TmRpc  *tmRpc.HTTP
}

func NewClient(cfg config.ConfigI, cdc *params.EncodingConfig, logger log.Logger) (*Client, error) {
	httpClient := &http.Client{}
	logger.Info("starting tm client...")
	tmRpcClient, err := tmRpc.NewWithClient(cfg.GetTendermintRPC(), "/websocket", httpClient)
	if err != nil {
		return nil, err
	}

	err = tmRpcClient.Start()
	if err != nil {
		logger.Error(
			"Tendermint client could not start",
			"address", cfg.GetTendermintRPC(),
			"error", err,
		)
		return nil, err
	}

	logger.Info("tm client started!")

	return &Client{
		TmRpc: tmRpcClient,
		cdc:   cdc,
	}, nil
}
