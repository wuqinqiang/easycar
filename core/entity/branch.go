package entity

import (
	"github.com/wuqinqiang/easycar/core/consts"
	"github.com/wuqinqiang/easycar/core/dao/gorm/model"
	"github.com/wuqinqiang/easycar/proto"
)

type (
	Branch struct {
		GId      string                 // global id
		BranchId string                 // branch id
		Url      string                 // branch request url (example grpc or http)
		ReqData  string                 // request data
		TranType consts.TransactionType // transaction type:tcc or saga or others
		PId      string                 // parent branch id
		Protocol string                 //http,grpc
		Action   consts.BranchAction    // action type of transaction
		State    consts.BranchState     // branch state
		//ChildrenList      []*Branch              //	children branch list
		EndTime int64 // end time
		// 07-10 add
		Level consts.Level // branch level in tree
	}
	BranchList []*Branch
)

func (b *Branch) IsSucceed() bool {
	return b.State == consts.BranchSucceed
}
func (b *Branch) IsBranchFailState() {

}

func (b *Branch) Convert() *model.Branch {
	return &model.Branch{
		GID:      b.GId,
		BranchID: b.BranchId,
		URL:      b.Url,
		ReqData:  b.ReqData,
		TranType: string(b.TranType),
		PID:      b.PId,
		Protocol: b.Protocol,
		Action:   string(b.Action),
		State:    string(b.State),
		EndTime:  int32(b.EndTime),
		Level:    int32(b.Level),
	}
}

// AssignmentByModel todo
func (b *Branch) AssignmentByModel(m *model.Branch) *Branch {
	b.GId = m.GID
	b.BranchId = m.BranchID
	b.Url = m.URL
	b.ReqData = m.ReqData
	b.TranType = consts.TransactionType(m.TranType)
	b.PId = m.PID
	b.Protocol = m.Protocol
	b.Action = consts.BranchAction(m.Action)
	b.State = consts.BranchState(m.State)
	b.EndTime = int64(m.EndTime)
	b.Level = consts.Level(m.Level)
	return b
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

func (list BranchList) Convert() []*model.Branch {
	var branches []*model.Branch
	for _, b := range list {
		branches = append(branches, b.Convert())
	}
	return branches
}

func (list BranchList) AssignmentByModel(mList []*model.Branch) BranchList {
	for i := range mList {
		var (
			b Branch
		)
		list = append(list, b.AssignmentByModel(mList[i]))
	}
	return list
}
func (list BranchList) AssignmentByGrpc(gid string, mList []*proto.RegisterReq_Branch) BranchList {
	for i := range mList {
		var (
			b Branch
		)
		b.GId = gid
		list = append(list, b.AssignmentByPb(mList[i]))
	}
	return list
}
