package gorm

import (
	"context"

	"github.com/wuqinqiang/easycar/conf"

	"github.com/wuqinqiang/easycar/tools"

	"github.com/wuqinqiang/easycar/core/entity"

	"github.com/wuqinqiang/easycar/core/consts"
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
		m entity.Global
	)
	m.GID = global.GetGId()
	err := g.query.Global.WithContext(ctx).Create(&m)
	return tools.WrapDbErr(err)
}

func (g GlobalImpl) GetGlobal(ctx context.Context, gid string) (entity.Global, error) {
	global := g.query.Global
	m, err := g.query.Global.WithContext(ctx).Where(global.GID.Eq(gid)).First()
	err = tools.WrapDbErr(err)
	if err != nil {
		return entity.Global{}, err
	}
	if m == nil {
		return entity.Global{}, nil
	}

	return *m, nil
}

func (g GlobalImpl) UpdateGlobalStateByGid(ctx context.Context, gid string,
	state consts.GlobalState) (int64, error) {
	global := g.query.Global
	result, err := g.query.Global.WithContext(ctx).
		Where(global.GID.Eq(gid)).Update(global.State, state)
	err = tools.WrapDbErr(err)
	return result.RowsAffected, err
}
