package service

import "github.com/wuqinqiang/easycar/internal/dao"

type RMOption func(rm *RM)

type RM struct {
	dao dao.TransactionDao
}

func NewRm(dao dao.TransactionDao) *RM {
	return &RM{dao: dao}
}

// Begin  begin a new transaction, return globleId
func (rm *RM) Begin() (gId string, err error) {
	panic("dc")
}

// Submit summit transaction
func (rm *RM) Submit(gId string) {
	panic("dd")
}

func (rm *RM) RegisterTccBranch() {

}

func (rm *RM) BranchRollBack() {

}
