package consts

import "github.com/wuqinqiang/easycar/proto"

func ConvertStateToGrpc(state GlobalState) proto.GlobalState {
	switch state {
	case Ready:
		return proto.GlobalState_Begin
	case Phase1Processing:
		return proto.GlobalState_Phase1Processing
	case Phase1Success:
		return proto.GlobalState_Phase1Success
	case Phase1Retrying:
		return proto.GlobalState_Phase1Retrying
	case Phase1Failed:
		return proto.GlobalState_Phase1Failed
	case Phase2Processing:
		return proto.GlobalState_Phase2Processing
	case Phase2Retrying:
		return proto.GlobalState_Phase2Retrying
	case Phase2Success:
		return proto.GlobalState_Phase2Success
	case Phase2Failed:
		return proto.GlobalState_Phase2Failed
	default:
	}
	return proto.GlobalState_GLOBAL_DEFAULT
}
