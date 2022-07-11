package entity

import (
	"github.com/wuqinqiang/easycar/core/consts"
	"github.com/wuqinqiang/easycar/core/dao/gorm/model"
)

type (
	Branch struct {
		GId               string                 // global id
		BranchId          string                 // branch id
		Url               string                 // branch request url (example grpc or http)
		ReqData           string                 // request data
		TranType          consts.TransactionType // transaction type:tcc or saga or others
		PId               string                 // parent branch id
		Protocol          string                 //http,grpc
		TransactionAction consts.BranchAction    // action type of transaction
		State             consts.BranchState     // branch state
		//ChildrenList      []*Branch              //	children branch list
		EndTime int64 // end time
		// 07-10 add
		Level consts.Level // branch level in tree
	}
	BranchList []*Branch
)

func (b *Branch) IsSucceed() bool {
	return b.State == consts.BranchSucceedState
}

func (b *Branch) Convert() *model.Branch {
	return &model.Branch{}
}

func (b *Branch) Assign(m *model.Branch) *Branch {
	return b
}

func (list BranchList) Convert() []*model.Branch {
	var branches []*model.Branch
	for _, b := range list {
		branches = append(branches, b.Convert())
	}
	return branches
}

func (list BranchList) Assign(mList []*model.Branch) BranchList {
	for _, b := range list {
		b.Assign(mList[0])
	}
	return list
}
