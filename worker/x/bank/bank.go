package bank

//import (
//	"github.com/Vitokz/eth_indexer/workers/indexer/types"
//	"github.com/Vitokz/eth_indexer/workers/indexer/x"
//	sdkTypes "github.com/cosmos/cosmos-sdk/types"
//	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
//	"github.com/pkg/errors"
//)
//
//// msgs types
//var (
//	msgSend      = sdkTypes.MsgTypeURL(&bankTypes.MsgSend{})
//	msgMultiSend = sdkTypes.MsgTypeURL(&bankTypes.MsgMultiSend{})
//)
//
//type Bank struct {
//	getHandlerFunc x.Handler
//}
//
//func NewBank(mode types.Mode) *Bank {
//	bank := Bank{}
//
//	switch mode {
//	case types.ModeBlockcout:
//		bank.getHandlerFunc = bank.getHandler
//	case types.ModeWQ:
//		bank.getHandlerFunc = bank.getHandlerForBlockScout
//	}
//
//	return &bank
//}
//
//func (s *Bank) GetHandler() types.HandlerI {
//	return s.getHandlerFunc()
//}
//
//func (s *Bank) Msgs() []string {
//	return []string{
//		msgSend,
//		msgMultiSend,
//	}
//}
//
//func (s *Bank) GetName() string {
//	return bankTypes.ModuleName
//}
//
//// WqMode handlers ------------------------------------------------
//
//func (s *Bank) getHandler() types.HandlerI {
//	return func(msg sdkTypes.Msg) error {
//		switch msg.(type) {
//		case *bankTypes.MsgSend:
//			return s.msgSend()
//		case *bankTypes.MsgMultiSend:
//			return s.msgMultiSend()
//		default:
//			return errors.Errorf("unrecognized %s message type: %T", s.GetName(), msg)
//		}
//	}
//}
//
//func (s *Bank) msgSend() error {
//
//	return nil
//}
//
//func (s *Bank) msgMultiSend() error {
//
//	return nil
//}
//
//// BlockScout handlers ------------------------------------------------
//
//func (s *Bank) getHandlerForBlockScout() types.HandlerI {
//	return func(msg sdkTypes.Msg) error {
//		switch msg.(type) {
//		case *bankTypes.MsgSend:
//			return s.msgSendForBlockScout()
//		case *bankTypes.MsgMultiSend:
//			return s.msgMultiSendForBlockScout()
//		default:
//			return errors.Errorf("unrecognized %s message type: %T", s.GetName(), msg)
//		}
//	}
//}
//
//func (s *Bank) msgSendForBlockScout() error {
//
//	return nil
//}
//
//func (s *Bank) msgMultiSendForBlockScout() error {
//
//	return nil
//}
