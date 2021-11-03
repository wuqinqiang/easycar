package dao

import (
	"context"
	"github.com/wuqinqiang/easycar/internal/model"
	"github.com/wuqinqiang/easycar/internal/query"
	"github.com/wuqinqiang/easycar/pkg/mysql"
	"github.com/wuqinqiang/easycar/pkg/utils"
)

type BranchState int32

const (
	BranchCommittedState BranchState = iota + 1
	BranchFinishedState
	BranchRollbackState
)

type BranchDao interface {
	CreateInBatches(ctx context.Context, branch []*model.Branch) error
	List(ctx context.Context, gid string) ([]*model.Branch, error)
	UpdateStateByGid(ctx context.Context, gid string, branchId string,
		state BranchState) (int64, error)
}

type BranchImpl struct {
	query *query.Query
}

func NewBranchImpl() BranchDao {
	return BranchImpl{query: query.Use(mysql.NewDb())}
}

func (g BranchImpl) CreateInBatches(ctx context.Context, branch []*model.Branch) error {
	err := g.query.Branch.WithContext(ctx).CreateInBatches(branch, len(branch))
	err = utils.WrapDbErr(err)
	return err
}

func (g BranchImpl) List(ctx context.Context, gid string) ([]*model.Branch, error) {
	branch := g.query.Branch
	list, err := g.query.Branch.WithContext(ctx).
		Where(branch.Gid.Eq(gid)).
		Find()
	err = utils.WrapDbErr(err)
	return list, err
}

func (g BranchImpl) UpdateStateByGid(ctx context.Context, gid string, branchId string,
	state BranchState) (int64, error) {
	branch := g.query.Branch
	result, err := g.query.Branch.WithContext(ctx).
		Where(branch.Gid.Eq(gid), branch.BranchID.Eq(branchId)).
		Update(branch.State, state)
	err = utils.WrapDbErr(err)
	return result.RowsAffected, err
}
