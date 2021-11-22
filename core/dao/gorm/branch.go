package gorm

import (
	"context"
	"strconv"
	"time"

	entity2 "github.com/wuqinqiang/easycar/core/entity"

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

func (g BranchImpl) CreateBatches(ctx context.Context, gId string, list []*entity2.Branch) error {
	var (
		branchList []*model.Branch
	)

	for i := range list {
		branchList = append(branchList, &model.Branch{
			Gid:        gId,
			URL:        list[i].GetUrl(),
			ReqData:    list[i].GetReqData(),
			BranchID:   gId + strconv.Itoa(i),
			BranchType: int32(list[i].GetBranchType()),
			State:      string(list[i].GetBranchState()),
			FinishTime: time.Now(),
			CreateTime: time.Time{},
			UpdateTime: time.Time{},
		})
	}

	err := g.query.Branch.WithContext(ctx).CreateInBatches(branchList, len(branchList))
	err = utils.WrapDbErr(err)
	return err
}

func (g BranchImpl) GetBranchList(ctx context.Context, gid string) ([]*entity2.Branch, error) {
	branch := g.query.Branch
	_, err := g.query.Branch.WithContext(ctx).
		Where(branch.Gid.Eq(gid)).
		Find()
	err = utils.WrapDbErr(err)
	return []*entity2.Branch{}, err
}

func (g BranchImpl) UpdateBranchStateByGid(ctx context.Context, gid string, state entity2.BranchState) (int64, error) {
	branch := g.query.Branch
	result, err := g.query.Branch.WithContext(ctx).
		Where(branch.Gid.Eq(gid)).
		Update(branch.State, state)
	err = utils.WrapDbErr(err)
	return result.RowsAffected, err
}
