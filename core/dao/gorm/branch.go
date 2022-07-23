package gorm

import (
	"context"

	"github.com/wuqinqiang/easycar/conf/common"

	"github.com/wuqinqiang/easycar/tools"

	"github.com/wuqinqiang/easycar/core/entity"

	"github.com/wuqinqiang/easycar/core/consts"

	"github.com/wuqinqiang/easycar/core/dao/gorm/query"
)

type BranchImpl struct {
	query *query.Query
}

func NewBranchImpl() BranchImpl {
	return BranchImpl{query: query.Use(common.GetDb())}
}

func (g BranchImpl) CreateBatches(ctx context.Context, list entity.BranchList) error {
	mList := list.Convert()
	err := g.query.Branch.WithContext(ctx).CreateInBatches(mList, len(mList))
	err = tools.WrapDbErr(err)
	return err
}

func (g BranchImpl) GetBranchList(ctx context.Context, gid string) (list entity.BranchList, err error) {
	q := g.query.Branch
	branches, err := g.query.Branch.WithContext(ctx).
		Where(q.GID.Eq(gid)).
		Find()
	if err = tools.WrapDbErr(err); err != nil {
		return
	}
	list = list.AssignmentByModel(branches)
	return
}

func (g BranchImpl) UpdateBranchStateByGid(ctx context.Context, gid string, state consts.BranchState) (int64, error) {
	branch := g.query.Branch
	result, err := g.query.Branch.WithContext(ctx).
		Where(branch.GID.Eq(gid)).
		Update(branch.State, state)
	err = tools.WrapDbErr(err)
	return result.RowsAffected, err
}
