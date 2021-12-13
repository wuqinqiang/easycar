package dao

import (
	"context"

	"github.com/wuqinqiang/easycar/pkg/entity"

	"github.com/wuqinqiang/easycar/pkg/common"
)

type TransactionDao interface {
	BranchDao
	GlobalDao
}

type BranchDao interface {
	CreateBatches(ctx context.Context, gId string, branch []*entity.Branch) error
	GetBranchList(ctx context.Context, gid string) ([]*entity.Branch, error)
	UpdateBranchStateByGid(ctx context.Context, gid string,
		state common.BranchState) (int64, error)
}

type GlobalDao interface {
	Create(ctx context.Context, global *entity.Global) (int32, error)
	First(ctx context.Context, gid string) (*entity.Global, error)
	UpdateGlobalStateByGid(ctx context.Context, gid string,
		state common.GlobalState) (int64, error)
}
