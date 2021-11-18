package service

import (
	"context"

	"github.com/wuqinqiang/easycar/internal/service/entity"
)

type TCC struct {
	*entity.Global
}

func (t *TCC) HandleBranches(ctx context.Context, branchList []*entity.Branch) error {
	return nil
}
