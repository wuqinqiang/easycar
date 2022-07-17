package dao

import "github.com/wuqinqiang/easycar/core/dao/gorm"

var (
	dao Dao
)

type Dao struct {
	BranchDao
	GlobalDao
}

func init() {
	dao = Dao{
		BranchDao: gorm.NewBranchImpl(),
		GlobalDao: gorm.NewGlobalImpl(),
	}
}

func GetTransaction() TransactionDao {
	return dao
}

func ReplaceGlobalDao(globalDao GlobalDao) {
	dao.GlobalDao = globalDao
}

func ReplaceBranchDao(branchDao BranchDao) {
	dao.BranchDao = branchDao
}
