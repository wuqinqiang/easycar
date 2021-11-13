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

type BranchImpl struct {
	query *query.Query
}

func NewBranchImpl() dao.BranchDao {
	return BranchImpl{query: query.Use(mysql.NewDb())}
}

func (g BranchImpl) CreateInBatches(ctx context.Context, branchEntities []*entity.Branch) error {
	var (
		branch []*model.Branch
	)
	err := g.query.Branch.WithContext(ctx).CreateInBatches(branch, len(branch))
	err = utils.WrapDbErr(err)
	return err
}

func (g BranchImpl) List(ctx context.Context, gid string) ([]*entity.Branch, error) {
	branch := g.query.Branch
	_, err := g.query.Branch.WithContext(ctx).
		Where(branch.Gid.Eq(gid)).
		Find()
	err = utils.WrapDbErr(err)
	return []*entity.Branch{}, err
}

func (g BranchImpl) UpdateBranchStateByGid(ctx context.Context, gid string, branchId string,
	state entity.BranchState) (int64, error) {
	branch := g.query.Branch
	result, err := g.query.Branch.WithContext(ctx).
		Where(branch.Gid.Eq(gid), branch.BranchID.Eq(branchId)).
		Update(branch.State, state)
	err = utils.WrapDbErr(err)
	return result.RowsAffected, err
}
