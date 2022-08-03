package dao

import (
	"context"
	"sync"

	"github.com/wuqinqiang/easycar/core/dao/gorm"

	"github.com/wuqinqiang/easycar/core/entity"

	"github.com/wuqinqiang/easycar/core/consts"
)

type TransactionDao interface {
	BranchDao
	GlobalDao
}

type BranchDao interface {
	CreateBatches(ctx context.Context, list entity.BranchList) error
	GetBranches(ctx context.Context, gid string) (entity.BranchList, error)
	UpdateBranchStateByGid(ctx context.Context, branchId string,
		state consts.BranchState, errMsg string) (int64, error)
}

type GlobalDao interface {
	CreateGlobal(ctx context.Context, global *entity.Global) error
	GetGlobal(ctx context.Context, gid string) (entity.Global, error)
	UpdateGlobalStateByGid(ctx context.Context, gid string,
		state consts.GlobalState) (int64, error)
}

var (
	dao  Dao
	once sync.Once
)

type Dao struct {
	BranchDao
	GlobalDao
}

func GetTransaction() TransactionDao {
	once.Do(func() {
		dao = Dao{
			BranchDao: gorm.NewBranchImpl(),
			GlobalDao: gorm.NewGlobalImpl(),
		}
	})
	return dao
}

func ReplaceGlobalDao(globalDao GlobalDao) {
	dao.GlobalDao = globalDao
}

func ReplaceBranchDao(branchDao BranchDao) {
	dao.BranchDao = branchDao
}
