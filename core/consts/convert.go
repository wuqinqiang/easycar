package consts

import "github.com/wuqinqiang/easycar/proto"

// ConvertStateToGrpc Convert global state to pb state
func ConvertStateToGrpc(state GlobalState) proto.GlobalState {
	switch state {
	case Init:
		return proto.GlobalState_INIT
	case Phase1Preparing:
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

// ConvertBranchStateToGrpc Convert branch state to pb state
func ConvertBranchStateToGrpc(state BranchState) proto.BranchState {
	switch state {
	case BranchInit:
		return proto.BranchState_B_INIT
	case BranchRetrying:
		return proto.BranchState_RETRYING
	case BranchSucceed:
		return proto.BranchState_SUCCEED
	case BranchFailState:
		return proto.BranchState_FAILED
	default:
	}
	return proto.BranchState_UN_KNOW_STATE
}

// ConvertBranchActionToGrpc Convert branch action to pb action
func ConvertBranchActionToGrpc(action BranchAction) proto.Action {
	switch action {
	case Try:
		return proto.Action_TRY
	case Confirm:
		return proto.Action_CONFIRM
	case Cancel:
		return proto.Action_CANCEL
	case Normal:
		return proto.Action_NORMAL
	case Compensation:
		return proto.Action_COMPENSATION
	default:
	}
	return proto.Action_UN_KNOW_TRANSACTION_TYPE
}

// ConvertTranTypeToGrpc Convert branch tranType to pb tranType
func ConvertTranTypeToGrpc(tranType TransactionType) proto.TranType {
	switch tranType {
	case TCC:
		return proto.TranType_TCC
	case SAGA:
		return proto.TranType_SAGE
	default:
	}
	return proto.TranType_UN_KNOW
}
