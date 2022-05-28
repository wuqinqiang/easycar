package entity

import "github.com/wuqinqiang/easycar/core/consts"

type Branch struct {
	GId               string
	Url               string // branch request url
	ReqData           string // request data
	RespData          string // resp data
	TranType          consts.TransactionType
	BranchId          string
	PId               string              // parent branch id
	Protocol          string              //http,grpc
	TransactionAction consts.BranchAction // action type of transaction
	State             consts.BranchState  // branch state
	ChildrenList      []*Branch           //	children branch list
	EndTime           int64               // end time
}
