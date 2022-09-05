package consts

import "github.com/wuqinqiang/easycar/proto"

func ConvertStateToGrpc(state GlobalState) proto.GlobalState {
	switch state {
	case Init:
		return proto.GlobalState_INIT
	case Phase1Processing:
		return proto.GlobalState_PHASE1_PROCESSING
	case Phase1Retrying:
		return proto.GlobalState_PHASE1_RETRYING
	case Phase1Failed:
		return proto.GlobalState_PHASE1_FAILED
	case Phase1Success:
		return proto.GlobalState_PHASE1_SUCCESS
	case Phase2Committing:
		return proto.GlobalState_PHASE2_COMMITTING
	case Phase2Rollbacking:
		return proto.GlobalState_PHASE2_ROLLBACKING
	case Phase2CommitFailed:
		return proto.GlobalState_PHASE2_COMMIT_FAILED
	case Phase2RollbackFailed:
		return proto.GlobalState_PHASE2_ROLLBACK_FAILED
	case Committed:
		return proto.GlobalState_COMMITTED
	case Rollbacked:
		return proto.GlobalState_ROLLBACKED
	default:
	}
	return proto.GlobalState_GLOBAL_DEFAULT
}
