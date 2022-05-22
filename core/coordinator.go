package core

import (
	"github.com/wuqinqiang/easycar/core/dao"
)

type CoordinatorInterface interface {
}

type Coordinator struct {
	// resty timeout
	restyTimeout int64 //second
	dao          dao.TransactionDao
}
