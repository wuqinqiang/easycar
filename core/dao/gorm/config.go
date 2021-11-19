package gorm

import (
	"context"

	"github.com/wuqinqiang/easycar/core/dao/gorm/query"
	"gorm.io/gorm"
)

// 暂时废弃
var GenManager *genManager

func InitGormDb(db *gorm.DB) {
	GenManager = &genManager{db}
}

type genManager struct {
	*gorm.DB
}

func (g *genManager) GetGormDb() *gorm.DB {
	if g.DB == nil {
		panic("init gorm err")
	}
	return g.DB
}

func (g *genManager) BeginTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return query.Use(g.GetGormDb()).Transaction(func(tx *query.Query) error {
		return fn(ctx)
	})
}
