package model

import "github.com/wuqinqiang/easycar/core/consts"

type Branch struct {
	gId               string
	url               string
	reqData           string
	branchId          string
	transactionAction consts.BranchAction
	state             consts.BranchState
}
