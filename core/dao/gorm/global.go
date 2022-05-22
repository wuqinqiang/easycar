package gorm

import (
	"context"

	"github.com/wuqinqiang/easycar/core/consts"

	"github.com/wuqinqiang/easycar/pkg/entity"

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
	first, err := g.query.Global.WithContext(ctx).
		Where(global.Gid.Eq(gid)).
		First()
	if err != nil {
		return nil, utils.WrapDbErr(err)
	}
	globalEntity := new(entity.Global)
	globalEntity.SetGId(first.Gid)
	globalEntity.SetTransactionName(consts.TransactionType(first.TransactionName))
	globalEntity.SetProtocol(first.Protocol)
	globalEntity.SetState(consts.GlobalState(first.State))
	return globalEntity, err
}

func (g GlobalImpl) UpdateGlobalStateByGid(ctx context.Context, gid string,
	state consts.GlobalState) (int64, error) {
	global := g.query.Global
	result, err := g.query.Global.WithContext(ctx).
		Where(global.Gid.Eq(gid)).Update(global.State, state)
	err = utils.WrapDbErr(err)
	return result.RowsAffected, err
}
