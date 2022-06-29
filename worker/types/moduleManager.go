package types

import (
	"context"
	"fmt"
	"github.com/Vitokz/lil_script/db"
	"github.com/Vitokz/lil_script/worker/x/staking"
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

// Module Manager ----------------------------------------------------------------------------------------

type ModuleManager struct {
	Modules []Module

	moduleHandlers map[string]HandlerI // map [' module name '] = this handler
	msgsHandlers   map[string]HandlerI // map[ ' msg type ' ] = this handler
}

func NewModuleManager() ModuleManager {
	var mm ModuleManager

	mm.Modules = []Module{
		staking.NewStaking(),
	}

	mm.msgsHandlers = make(map[string]HandlerI)
	mm.moduleHandlers = make(map[string]HandlerI)

	mm.setMsgsAndThemHandlers()
	return mm
}

// GetMsgHandler searches for the required handler for the given message type
func (m *ModuleManager) GetMsgHandler(msg string) (HandlerI, error) {
	handler, ok := m.msgsHandlers[msg]
	if ok {
		return handler, nil
	} else {
		return nil, fmt.Errorf("not found handler for msg type: %s", msg)
	}
}

// setMsgsAndThemHandlers fills in manager data
func (m *ModuleManager) setMsgsAndThemHandlers() {
	for _, module := range m.Modules {
		handler := module.GetHandler()
		for _, msg := range module.Msgs() {
			m.msgsHandlers[msg] = handler
		}

		m.moduleHandlers[module.GetName()] = handler
	}
}
