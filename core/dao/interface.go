package dao

import (
	"context"

	"github.com/wuqinqiang/easycar/core/entity"

	"github.com/wuqinqiang/easycar/core/consts"
)

type TransactionDao interface {
	BranchDao
	GlobalDao
}

type BranchDao interface {
	CreateBatches(ctx context.Context, list entity.BranchList) error
	GetBranchList(ctx context.Context, gid string) (entity.BranchList, error)
	UpdateBranchStateByGid(ctx context.Context, gid string,
		state consts.BranchState) (int64, error)
}

type GlobalDao interface {
	CreateGlobal(ctx context.Context, global *entity.Global) error
	GetGlobal(ctx context.Context, gid string) (*entity.Global, error)
	UpdateGlobalStateByGid(ctx context.Context, gid string,
		state consts.GlobalState) (int64, error)
}
