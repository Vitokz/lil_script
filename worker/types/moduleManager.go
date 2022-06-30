package types

import (
	"fmt"
	"github.com/Vitokz/lil_script/worker/x"
	"github.com/Vitokz/lil_script/worker/x/staking"
)

// Module Manager ----------------------------------------------------------------------------------------

type ModuleManager struct {
	Modules []x.Module

	moduleHandlers map[string]x.HandlerI // map [' module name '] = this handler
	msgsHandlers   map[string]x.HandlerI // map[ ' msg type ' ] = this handler
}

func NewModuleManager() ModuleManager {
	var mm ModuleManager

	mm.Modules = []x.Module{
		staking.NewStaking(),
	}

	mm.msgsHandlers = make(map[string]x.HandlerI)
	mm.moduleHandlers = make(map[string]x.HandlerI)

	mm.setMsgsAndThemHandlers()
	return mm
}

// GetMsgHandler searches for the required handler for the given message type
func (m *ModuleManager) GetMsgHandler(msg string) (x.HandlerI, error) {
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
