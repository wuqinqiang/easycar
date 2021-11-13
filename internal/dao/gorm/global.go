package gorm

import (
	"context"
	"github.com/wuqinqiang/easycar/internal/dao"
	"github.com/wuqinqiang/easycar/internal/gorm/model"
	"github.com/wuqinqiang/easycar/internal/gorm/query"
	"github.com/wuqinqiang/easycar/internal/service/entity"
	"github.com/wuqinqiang/easycar/pkg/mysql"
	"github.com/wuqinqiang/easycar/pkg/utils"
)

type GlobalImpl struct {
	query *query.Query
}

func NewGlobalImpl() dao.GlobalDao {
	return GlobalImpl{query: query.Use(mysql.NewDb())}
}

func (g GlobalImpl) Create(ctx context.Context, globalEntity *entity.Global) (int32, error) {
	var (
		global model.Global
	)
	err := g.query.Global.WithContext(ctx).Create(&global)
	err = utils.WrapDbErr(err)
	return global.ID, err
}

func (g GlobalImpl) First(ctx context.Context, gid string) (*entity.Global, error) {
	global := g.query.Global
	_, err := g.query.Global.WithContext(ctx).
		Where(global.Gid.Eq(gid)).
		First()
	err = utils.WrapDbErr(err)
	return &entity.Global{}, err
}

func (g GlobalImpl) UpdateGlobalStateByGid(ctx context.Context, gid string,
	state entity.GlobalState) (int64, error) {
	global := g.query.Global
	result, err := g.query.Global.WithContext(ctx).
		Where(global.Gid.Eq(gid)).Update(global.State, state)
	err = utils.WrapDbErr(err)
	return result.RowsAffected, err
}
