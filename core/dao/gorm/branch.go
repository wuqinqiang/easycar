package gorm

import (
	"context"

	"github.com/wuqinqiang/easycar/core/entity"

	"github.com/wuqinqiang/easycar/core/consts"

	"github.com/wuqinqiang/easycar/core/dao/gorm/query"
	"github.com/wuqinqiang/easycar/pkg/mysql"
	"github.com/wuqinqiang/easycar/pkg/utils"
)

type BranchImpl struct {
	query *query.Query
}

func NewBranchImpl() BranchImpl {
	return BranchImpl{query: query.Use(mysql.NewDb())}
}

func (g BranchImpl) CreateBatches(ctx context.Context, list entity.BranchList) error {
	mList := list.Convert()
	err := g.query.Branch.WithContext(ctx).CreateInBatches(mList, len(mList))
	err = utils.WrapDbErr(err)
	return err
}

func (g BranchImpl) GetBranchList(ctx context.Context, gid string) (list entity.BranchList, err error) {
	q := g.query.Branch
	branches, err := g.query.Branch.WithContext(ctx).
		Where(q.Gid.Eq(gid)).
		Find()
	if err = utils.WrapDbErr(err); err != nil {
		return
	}
	list = list.Assign(branches)
	return
}

func (g BranchImpl) UpdateBranchStateByGid(ctx context.Context, gid string, state consts.BranchState) (int64, error) {
	branch := g.query.Branch
	result, err := g.query.Branch.WithContext(ctx).
		Where(branch.Gid.Eq(gid)).
		Update(branch.State, state)
	err = utils.WrapDbErr(err)
	return result.RowsAffected, err
}
