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

func (b *Branch) Convert() *model.Branch {
	return &model.Branch{}
}

func (list BranchList) Convert() []*model.Branch {
	var branches []*model.Branch
	for _, b := range list {
		branches = append(branches, b.Convert())
	}
	return branches
}
