package entity

import (
	"github.com/wuqinqiang/easycar/core/consts"
	"github.com/wuqinqiang/easycar/proto"
)

func GetTranTypeByPb(protoType proto.TranType) consts.TransactionType {
	switch protoType {
	case proto.TranType_TCC:
		return consts.TCC
	case proto.TranType_SAGE:
		return consts.SAGA
	default:
	}
	return consts.TransactionUnknown
}

func GetActionByPb(action proto.Action) consts.BranchAction {
	switch action {
	case proto.Action_Try:
		return consts.Try
	case proto.Action_Confirm:
		return consts.Confirm
	case proto.Action_Cancel:
		return consts.Cancel
	case proto.Action_Normal:
		return consts.Normal
	case proto.Action_Compensation:
		return consts.Compensation
	default:
	}
	return consts.ActionUnknown
}

// AssignmentByPb todo
func (b *Branch) AssignmentByPb(m *proto.RegisterReq_Branch) *Branch {
	b.Url = m.GetUri()
	b.ReqData = m.GetReqData()
	b.TranType = GetTranTypeByPb(m.GetTranType())
	b.Protocol = m.GetProtocol()
	b.Action = GetActionByPb(m.GetAction())
	b.Level = consts.Level(m.GetLevel())
	return b
}

func (list BranchList) AssignmentByGrpc(gid string, mList []*proto.RegisterReq_Branch) BranchList {
	for i := range mList {
		var (
			b Branch
		)
		b.GID = gid
		list = append(list, b.AssignmentByPb(mList[i]))
	}
	return list
}
