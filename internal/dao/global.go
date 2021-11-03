package dao

import (
	"context"
	"github.com/wuqinqiang/easycar/internal/model"
	"github.com/wuqinqiang/easycar/internal/query"
	"github.com/wuqinqiang/easycar/pkg/mysql"
	"github.com/wuqinqiang/easycar/pkg/utils"
)

type GlobalState int32

const (
	PreparedState GlobalState = iota + 1
	SubmittedState
	AbortingState
	RollbackState
)

type GlobalDao interface {
	Create(ctx context.Context, global *model.Global) (int32, error)
	First(ctx context.Context, gid string) (*model.Global, error)
	UpdateStateByGid(ctx context.Context, gid string,
		state GlobalState) (int64, error)
}

type GlobalImpl struct {
	query *query.Query
}

func NewGlobalImpl() GlobalDao {
	return GlobalImpl{query: query.Use(mysql.NewDb())}
}

func (g GlobalImpl) Create(ctx context.Context, global *model.Global) (int32, error) {
	err := g.query.Global.WithContext(ctx).Create(global)
	err = utils.WrapDbErr(err)
	return global.ID, err
}

func (g GlobalImpl) First(ctx context.Context, gid string) (*model.Global, error) {
	global := g.query.Global
	first, err := g.query.Global.WithContext(ctx).
		Where(global.Gid.Eq(gid)).
		First()
	err = utils.WrapDbErr(err)
	return first, err
}

func (g GlobalImpl) UpdateStateByGid(ctx context.Context, gid string,
	state GlobalState) (int64, error) {
	global := g.query.Global
	result, err := g.query.Global.WithContext(ctx).
		Where(global.Gid.Eq(gid)).Update(global.State, state)
	err = utils.WrapDbErr(err)
	return result.RowsAffected, err
}
