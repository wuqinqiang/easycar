package core

import (
	"context"

	"github.com/panjf2000/ants/v2"
)

type Pool struct {
	pool *ants.Pool
	opts opts
}

func NewPool(optFns ...OptFn) (*Pool, error) {
	opts := defaultOpts
	for _, fn := range optFns {
		fn(&opts)
	}
	pool, err := ants.NewPool(int(opts.poolSize))
	if err != nil {
		return nil, err
	}
	return &Pool{pool: pool, opts: opts}, nil
}

func (p *Pool) executeTask(ctx context.Context, task task) error {
	return p.pool.Submit(func() {
	})
}
