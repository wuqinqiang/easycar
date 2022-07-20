package dao

import (
	"sync"

	"github.com/wuqinqiang/easycar/core/dao/gorm"
)

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
