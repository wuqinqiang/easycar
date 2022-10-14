package etcdx

import (
	"context"
	"time"
)

type OptFn func(options *Options)

type Options struct {
	ttl time.Duration
	ctx context.Context
}

func newDefault() Options {
	return Options{ttl: 10 * time.Second, ctx: context.Background()}
}
