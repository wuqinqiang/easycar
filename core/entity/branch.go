package entity

import (
	"github.com/wuqinqiang/easycar/core/consts"
	"github.com/wuqinqiang/easycar/proto"
)

type (
	Branch struct {
		GID      string                 `gorm:"column:g_id;type:varchar(255);not null"`                      // global id
		BranchId string                 `gorm:"column:branch_id;type:varchar(255);not null"`                 // branch id
		Url      string                 `gorm:"column:url;type:varchar(255);not null"`                       // branch request url (example grpc or http)
		ReqData  string                 `gorm:"column:req_data;type:varchar(255);not null"`                  // request data
		TranType consts.TransactionType `gorm:"column:tran_type;type:varchar(255);not null"`                 // transaction type:tcc or saga or others
		PId      string                 `gorm:"column:p_id;type:varchar(255);not null"`                      // parent branch id
		Protocol string                 `gorm:"column:protocol;type:varchar(255);not null;default:http"`     //http,grpc
		Action   consts.BranchAction    `gorm:"column:action;type:varchar(255);not null"`                    // action type of transaction
		State    consts.BranchState     `gorm:"column:State;type:varchar(255);not null;default:branchReady"` // branch State
		//ChildrenList      []*Branch               //	children branch list
		EndTime int64 `gorm:"column:end_time;type:int;not null;default:0"`
		// 07-10 add
		Level consts.Level `gorm:"column:level;type:int;not null;default:1"` // branch level in tree
	}
	BranchList []*Branch
)

func (b *Branch) IsSucceed() bool {
	return b.State == consts.BranchSucceed
}
func (b *Branch) IsBranchFailState() {

}

// AssignmentByPb todo
func (b *Branch) AssignmentByPb(m *proto.RegisterReq_Branch) *Branch {
	b.BranchId = m.GetBranchId()
	b.Url = m.GetUri()
	b.ReqData = m.GetReqData()
	b.TranType = consts.TransactionType(m.GetTranType().String())
	b.PId = m.GetPid()
	b.Protocol = m.GetProtocol()
	b.Action = consts.BranchAction(m.GetAction().String())
	b.State = consts.BranchState(m.GetState().String())
	b.Level = consts.Level(m.GetLevel())
	return b
}

func (b *Branch) IsTcc() bool {
	return b.TranType == consts.TCC
}
func (b *Branch) IsTccTry() bool {
	return b.Action == consts.Try && b.IsTcc()
}

func (b *Branch) IsSAGA() bool {
	return b.TranType == consts.SAGA
}

func (b *Branch) IsSAGANormal() bool {
	return b.Action == consts.Normal && b.IsSAGA()
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
