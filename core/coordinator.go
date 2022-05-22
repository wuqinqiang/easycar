package core

import (
	"context"

	entity2 "github.com/wuqinqiang/easycar/pkg/entity"

	"github.com/wuqinqiang/easycar/core/dao"
)

type CoordinatorInterface interface {
	Begin(ctx context.Context, entity *entity2.Global) (gId string, err error)
	BranchRegister(ctx context.Context, branch *entity2.Branch) (state string, err error)
	Commit(ctx context.Context, gId string) (state string, err error)
	RollBack(ctx context.Context, gId string) (state string, err error)
}

type Coordinator struct {
	// resty timeout
	restyTimeout int64 //second
	dao          dao.TransactionDao
}
