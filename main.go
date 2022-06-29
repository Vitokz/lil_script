package main

import (
	"github.com/Vitokz/lil_script/config"
	"github.com/Vitokz/lil_script/rpc"
	"github.com/Vitokz/lil_script/worker"
	tmcfg "github.com/tendermint/tendermint/config"
	tmflags "github.com/tendermint/tendermint/libs/cli/flags"
	"github.com/tendermint/tendermint/libs/log"
	ethApp "github.com/tharsis/ethermint/app"
	ethEncoding "github.com/tharsis/ethermint/encoding"
	"os"
)

func main() {
	cfg := config.ParseConfigInConfigFile()

	cdc := ethEncoding.MakeConfig(ethApp.ModuleBasics)

	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	logger, _ = tmflags.ParseLogLevel("info", logger, tmcfg.DefaultLogLevel)

	cm, err := rpc.NewClient(cfg, &cdc, logger)
	if err != nil {
		panic(err)
	}
	_ = cm

	w, err := worker.NewWorker(cfg, &cdc, logger)
	if err != nil {
		panic(err)
	}

	w.StartWorker(worker.TxEvent)
}
