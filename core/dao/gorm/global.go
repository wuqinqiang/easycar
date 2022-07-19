package gorm

import (
	"context"

	"github.com/wuqinqiang/easycar/conf"
	"github.com/wuqinqiang/easycar/tools"

	"github.com/wuqinqiang/easycar/core/entity"

	"github.com/wuqinqiang/easycar/core/consts"
	"github.com/wuqinqiang/easycar/core/dao/gorm/model"
	"github.com/wuqinqiang/easycar/core/dao/gorm/query"
)

type GlobalImpl struct {
	query *query.Query
}

func NewGlobalImpl() GlobalImpl {
	return GlobalImpl{query: query.Use(conf.GetDb())}
}

func (g GlobalImpl) CreateGlobal(ctx context.Context, global *entity.Global) error {
	var (
		m model.Global
	)
	m.Gid = global.GetGId()
	err := g.query.Global.WithContext(ctx).Create(&m)
	return tools.WrapDbErr(err)
}

func (g GlobalImpl) GetGlobal(ctx context.Context, gid string) (*entity.Global, error) {
	return nil, nil
}

func (g GlobalImpl) UpdateGlobalStateByGid(ctx context.Context, gid string,
	state consts.GlobalState) (int64, error) {
	global := g.query.Global
	result, err := g.query.Global.WithContext(ctx).
		Where(global.Gid.Eq(gid)).Update(global.State, state)
	err = tools.WrapDbErr(err)
	return result.RowsAffected, err
}
