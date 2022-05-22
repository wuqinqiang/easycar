package gorm

import (
	"context"

	"github.com/wuqinqiang/easycar/core/consts"

	"github.com/wuqinqiang/easycar/core/dao"
	"github.com/wuqinqiang/easycar/core/dao/gorm/model"
	"github.com/wuqinqiang/easycar/core/dao/gorm/query"
	"github.com/wuqinqiang/easycar/pkg/mysql"
	"github.com/wuqinqiang/easycar/pkg/utils"
)

type BranchImpl struct {
	query *query.Query
}

func NewBranchImpl() dao.BranchDao {
	return BranchImpl{query: query.Use(mysql.NewDb())}
}

func (g BranchImpl) CreateBatches(ctx context.Context, list []*model.Branch) error {
	err := g.query.Branch.WithContext(ctx).CreateInBatches(list, len(list))
	err = utils.WrapDbErr(err)
	return err
}

func (g BranchImpl) GetBranchList(ctx context.Context, gid string) ([]*model.Branch, error) {
	branch := g.query.Branch
	_, err := g.query.Branch.WithContext(ctx).
		Where(branch.Gid.Eq(gid)).
		Find()
	err = utils.WrapDbErr(err)
	return []*model.Branch{}, err
}

func (g BranchImpl) UpdateBranchStateByGid(ctx context.Context, gid string, state consts.BranchState) (int64, error) {
	branch := g.query.Branch
	result, err := g.query.Branch.WithContext(ctx).
		Where(branch.Gid.Eq(gid)).
		Update(branch.State, state)
	err = utils.WrapDbErr(err)
	return result.RowsAffected, err
}
