package gorm

import (
	"context"

	entity2 "github.com/wuqinqiang/easycar/core/entity"

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

func (g GlobalImpl) Create(ctx context.Context, globalEntity *entity2.Global) (int32, error) {
	var (
		global model.Global
	)
	err := g.query.Global.WithContext(ctx).Create(&global)
	err = utils.WrapDbErr(err)
	return global.ID, err
}

func (g GlobalImpl) First(ctx context.Context, gid string) (*entity2.Global, error) {
	global := g.query.Global
	first, err := g.query.Global.WithContext(ctx).
		Where(global.Gid.Eq(gid)).
		First()
	if err != nil {
		return nil, utils.WrapDbErr(err)
	}
	globalEntity := new(entity2.Global)
	globalEntity.SetGId(first.Gid)
	globalEntity.SetTransactionName(entity2.TransactionName(first.TransactionName))
	globalEntity.SetProtocol(first.Protocol)
	globalEntity.SetState(entity2.GlobalState(first.State))
	return globalEntity, err
}

func (g GlobalImpl) UpdateGlobalStateByGid(ctx context.Context, gid string,
	state entity2.GlobalState) (int64, error) {
	global := g.query.Global
	result, err := g.query.Global.WithContext(ctx).
		Where(global.Gid.Eq(gid)).Update(global.State, state)
	err = utils.WrapDbErr(err)
	return result.RowsAffected, err
}
