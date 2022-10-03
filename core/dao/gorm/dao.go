package gorm

import (
	"github.com/wuqinqiang/easycar/core/dao"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type Dao struct {
	dao.BranchDao
	dao.GlobalDao
}

func NewDao(gorm *gorm.DB) Dao {
	db = gorm
	return Dao{
		BranchDao: NewBranchImpl(),
		GlobalDao: NewGlobalImpl(),
	}
}
