package coordinator

import (
	"context"

	"github.com/wuqinqiang/easycar/core/entity"
)

type Executor interface {
	Phase1(ctx context.Context, global *entity.Global, branches entity.BranchList) error
	Phase2(ctx context.Context, global *entity.Global, branches entity.BranchList) error

	// Close when the server stop
	Close(ctx context.Context) error
}
