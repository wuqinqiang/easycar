package dao

import (
	"context"

	"github.com/wuqinqiang/easycar/core/dao/gorm/model"

	"github.com/wuqinqiang/easycar/core/consts"
)

type TransactionDao interface {
	BranchDao
	GlobalDao
}

type BranchDao interface {
	CreateBatches(ctx context.Context, branch []*model.Branch) error
	GetBranchList(ctx context.Context, gid string) ([]*model.Branch, error)
	UpdateBranchStateByGid(ctx context.Context, gid string,
		state consts.BranchState) (int64, error)
}

type GlobalDao interface {
	Create(ctx context.Context, global *model.Global) (int32, error)
	First(ctx context.Context, gid string) (*model.Global, error)
	UpdateGlobalStateByGid(ctx context.Context, gid string,
		state consts.GlobalState) (int64, error)
}
