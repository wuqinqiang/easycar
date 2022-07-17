package consts

import "github.com/wuqinqiang/easycar/proto"

func ConvertStateToGrpc(state GlobalState) proto.GlobalState {
	switch state {
	case Begin:
		return proto.GlobalState_Begin
	case GlobalCommitting:
		return proto.GlobalState_GlobalCommitting
	case GlobalCommitted:
		return proto.GlobalState_GlobalCommitted
	case GlobalCommitRetrying:
		return proto.GlobalState_GlobalCommitRetrying
	case GlobalCommitFailed:
		return proto.GlobalState_GlobalCommitFailed
	case GlobalRollBacking:
		return proto.GlobalState_GlobalRollBacking
	case GlobalRollBacked:
		return proto.GlobalState_GlobalRollBacked
	case GlobalRollBackRetrying:
		return proto.GlobalState_GlobalRollBackRetrying
	case GlobalRollBackFailed:
		return proto.GlobalState_GlobalRollBackFailed
	default:
	}
	return proto.GlobalState_GLOBAL_DEFAULT
}
