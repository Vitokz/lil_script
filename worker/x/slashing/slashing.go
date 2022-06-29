package slashing

//import (
//	"github.com/Vitokz/eth_indexer/workers/indexer/types"
//	"github.com/Vitokz/eth_indexer/workers/indexer/x"
//	sdkTypes "github.com/cosmos/cosmos-sdk/types"
//	slashingTypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
//	"github.com/pkg/errors"
//)
//
//// msgs types
//var (
//	msgUnjail = sdkTypes.MsgTypeURL(&slashingTypes.MsgUnjail{})
//)
//
//type Slashing struct {
//	getHandlerFunc x.Handler
//}
//
//func NewSlashing(mode types.Mode) *Slashing {
//	slashing := Slashing{}
//
//	switch mode {
//	case types.ModeBlockcout:
//		slashing.getHandlerFunc = slashing.getHandler
//	case types.ModeWQ:
//		slashing.getHandlerFunc = slashing.getHandlerForBlockScout
//	}
//
//	return &slashing
//}
//
//func (s *Slashing) GetHandler() types.HandlerI {
//	return s.getHandlerFunc()
//}
//
//func (s *Slashing) Msgs() []string {
//	return []string{
//		msgUnjail,
//	}
//}
//
//func (s *Slashing) GetName() string {
//	return slashingTypes.ModuleName
//}
//
//// WqMode handlers ------------------------------------------------
//
//func (s *Slashing) getHandler() types.HandlerI {
//	return func(msg sdkTypes.Msg) error {
//		switch msg.(type) {
//		case *slashingTypes.MsgUnjail:
//			return s.msgUnjail()
//		default:
//			return errors.Errorf("unrecognized %s message type: %T", s.GetName(), msg)
//		}
//	}
//}
//
//func (s *Slashing) msgUnjail() error {
//
//	return nil
//}
//
//// BlockScout handlers ------------------------------------------------
//
//func (s *Slashing) getHandlerForBlockScout() types.HandlerI {
//	return func(msg sdkTypes.Msg) error {
//		switch msg.(type) {
//		case *slashingTypes.MsgUnjail:
//			return s.msgUnjailForBlockScout()
//		default:
//			return errors.Errorf("unrecognized %s message type: %T", s.GetName(), msg)
//		}
//	}
//}
//
//func (s *Slashing) msgUnjailForBlockScout() error {
//
//	return nil
//}
