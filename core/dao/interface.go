package dao

import (
	"context"

	entity2 "github.com/wuqinqiang/easycar/core/entity"
)

type TransactionDao interface {
	BranchDao
	GlobalDao
}

type BranchDao interface {
	CreateBatches(ctx context.Context, gId string, branch []*entity2.Branch) error
	GetBranchList(ctx context.Context, gid string) ([]*entity2.Branch, error)
	UpdateBranchStateByGid(ctx context.Context, gid string,
		state entity2.BranchState) (int64, error)
}

type GlobalDao interface {
	Create(ctx context.Context, global *entity2.Global) (int32, error)
	First(ctx context.Context, gid string) (*entity2.Global, error)
	UpdateGlobalStateByGid(ctx context.Context, gid string,
		state entity2.GlobalState) (int64, error)
}
