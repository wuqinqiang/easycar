package gorm

import (
	"context"

	"github.com/wuqinqiang/easycar/tools"

	"github.com/wuqinqiang/easycar/core/entity"

	"github.com/wuqinqiang/easycar/core/consts"

	"github.com/wuqinqiang/easycar/core/dao/gorm/query"
)

type BranchImpl struct {
	query *query.Query
}

func NewBranchImpl() BranchImpl {
	return BranchImpl{query: query.Use(db)}
}

func (g BranchImpl) CreateBatches(ctx context.Context, list entity.BranchList) error {
	err := g.query.Branch.WithContext(ctx).CreateInBatches(list, len(list))
	err = tools.WrapDbErr(err)
	return err
}

func (g BranchImpl) GetBranches(ctx context.Context, gid string) (list entity.BranchList, err error) {
	q := g.query.Branch
	list, err = g.query.Branch.WithContext(ctx).
		Where(q.GID.Eq(gid)).
		Find()
	if err = tools.WrapDbErr(err); err != nil {
		return
	}
	return
}

func (g BranchImpl) UpdateBranchStateByGid(ctx context.Context, branchId string, state consts.BranchState, errmsg string) (int64, error) {
	branch := g.query.Branch
	result, err := g.query.Branch.WithContext(ctx).
		Where(branch.BranchId.Eq(branchId)).
		UpdateSimple(branch.State.Value(string(state)), branch.LastErrMsg.Value(errmsg))
	err = tools.WrapDbErr(err)
	return result.RowsAffected, err
}
