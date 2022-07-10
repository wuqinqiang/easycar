package gorm

import (
	"context"

	"github.com/wuqinqiang/easycar/core/entity"

	"github.com/wuqinqiang/easycar/core/consts"
	"github.com/wuqinqiang/easycar/core/dao"
	"github.com/wuqinqiang/easycar/core/dao/gorm/model"
	"github.com/wuqinqiang/easycar/core/dao/gorm/query"
	"github.com/wuqinqiang/easycar/pkg/mysql"
	"github.com/wuqinqiang/easycar/pkg/utils"
)

type GlobalImpl struct {
	query *query.Query
}

func NewGlobalImpl() dao.GlobalDao {
	return GlobalImpl{query: query.Use(mysql.NewDb())}
}

func (g GlobalImpl) CreateGlobal(ctx context.Context, global *entity.Global) error {
	var (
		m model.Global
	)
	m.Gid = global.GetGId()
	err := g.query.Global.WithContext(ctx).Create(&m)
	return utils.WrapDbErr(err)
}

func (g GlobalImpl) GetGlobal(ctx context.Context, gid string) (*entity.Global, error) {
	return nil, nil
}

func (g GlobalImpl) UpdateGlobalStateByGid(ctx context.Context, gid string,
	state consts.GlobalState) (int64, error) {
	global := g.query.Global
	result, err := g.query.Global.WithContext(ctx).
		Where(global.Gid.Eq(gid)).Update(global.State, state)
	err = utils.WrapDbErr(err)
	return result.RowsAffected, err
}
