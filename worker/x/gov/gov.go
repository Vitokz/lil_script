package gov

//import (
//	"github.com/Vitokz/eth_indexer/workers/indexer/types"
//	"github.com/Vitokz/eth_indexer/workers/indexer/x"
//	sdkTypes "github.com/cosmos/cosmos-sdk/types"
//	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
//	"github.com/pkg/errors"
//)
//
//// msgs types
//var (
//	msgVote           = sdkTypes.MsgTypeURL(&govTypes.MsgVote{})
//	msgDeposit        = sdkTypes.MsgTypeURL(&govTypes.MsgDeposit{})
//	msgVoteWeighted   = sdkTypes.MsgTypeURL(&govTypes.MsgVoteWeighted{})
//	msgSubmitProposal = sdkTypes.MsgTypeURL(&govTypes.MsgSubmitProposal{})
//)
//
//type Gov struct {
//	getHandlerFunc x.Handler
//}
//
//func NewGov(mode types.Mode) *Gov {
//	gov := Gov{}
//
//	switch mode {
//	case types.ModeBlockcout:
//		gov.getHandlerFunc = gov.getHandler
//	case types.ModeWQ:
//		gov.getHandlerFunc = gov.getHandlerForBlockScout
//	}
//
//	return &gov
//}
//
//func (s *Gov) GetHandler() types.HandlerI {
//	return s.getHandlerFunc()
//}
//
//func (s *Gov) Msgs() []string {
//	return []string{
//		msgVote,
//		msgDeposit,
//		msgSubmitProposal,
//		msgVoteWeighted,
//	}
//}
//
//func (s *Gov) GetName() string {
//	return govTypes.ModuleName
//}
//
//// WqMode handlers ------------------------------------------------
//
//func (s *Gov) getHandler() types.HandlerI {
//	return func(msg sdkTypes.Msg) error {
//		switch msg.(type) {
//		case *govTypes.MsgDeposit:
//			return s.msgDeopist()
//		case *govTypes.MsgSubmitProposal:
//			return s.msgSubmitProposal()
//		case *govTypes.MsgVote:
//			return s.msgVote()
//		case *govTypes.MsgVoteWeighted:
//			return s.msgVoteWeighted()
//		default:
//			return errors.Errorf("unrecognized %s message type: %T", s.GetName(), msg)
//		}
//	}
//}
//
//func (s *Gov) msgVote() error {
//
//	return nil
//}
//
//func (s *Gov) msgDeopist() error {
//
//	return nil
//}
//
//func (s *Gov) msgVoteWeighted() error {
//
//	return nil
//}
//
//func (s *Gov) msgSubmitProposal() error {
//
//	return nil
//}
//
//// BlockScout handlers ------------------------------------------------
//
//func (s *Gov) getHandlerForBlockScout() types.HandlerI {
//	return func(msg sdkTypes.Msg) error {
//		switch msg.(type) {
//		case *govTypes.MsgDeposit:
//			return s.msgDeopistForBlockScout()
//		case *govTypes.MsgSubmitProposal:
//			return s.msgSubmitProposalForBlockScout()
//		case *govTypes.MsgVote:
//			return s.msgVoteForBlockScout()
//		case *govTypes.MsgVoteWeighted:
//			return s.msgVoteWeightedForBlockScout()
//		default:
//			return errors.Errorf("unrecognized %s message type: %T", s.GetName(), msg)
//		}
//	}
//}
//
//func (s *Gov) msgVoteForBlockScout() error {
//
//	return nil
//}
//
//func (s *Gov) msgDeopistForBlockScout() error {
//
//	return nil
//}
//
//func (s *Gov) msgVoteWeightedForBlockScout() error {
//
//	return nil
//}
//
//func (s *Gov) msgSubmitProposalForBlockScout() error {
//
//	return nil
//}
