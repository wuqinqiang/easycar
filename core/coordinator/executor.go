package coordinator

import "context"

type Executor interface {
	Execute(ctx context.Context) error
}
