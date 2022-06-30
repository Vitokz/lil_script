package x

import (
	"context"
	"github.com/Vitokz/lil_script/db"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	web3 "github.com/ethereum/go-ethereum/ethclient"
	"github.com/tendermint/tendermint/libs/log"
)

// Module is interfaces for modules to be accepted by the manager
type Module interface {
	GetHandler() HandlerI
	Msgs() []string
	GetName() string
}

type HandlerI func(ctx context.Context, web3 *web3.Client, db *db.DB, logger log.Logger, msg sdkTypes.Msg) error
