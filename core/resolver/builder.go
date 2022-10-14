package resolver

import (
	"context"
	"strings"
	"time"

	"github.com/wuqinqiang/easycar/tools"

	"github.com/wuqinqiang/easycar/core/registry"
	"google.golang.org/grpc/resolver"
)

var _ resolver.Builder = (*builder)(nil)

type builder struct {
	discoverer   registry.Discovery
	watchTimeout time.Duration
}

func NewBuilder(d registry.Discovery) *builder {
	return &builder{
		discoverer:   d,
		watchTimeout: time.Second * 8,
	}
}

func (b builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	ctx, cancel := context.WithCancel(context.Background())
	watchCtx, watchCancel := context.WithTimeout(ctx, b.watchTimeout)
	defer watchCancel()
	watch, err := b.discoverer.Watch(watchCtx, strings.TrimPrefix(target.URL.Path, "/"))
	if err != nil {
		cancel()
		return nil, err
	}
	r := NewDefaultResolver(ctx, cancel, cc, watch)
	tools.GoSafe(func() {
		r.watch()
	})
	return r, nil
}

func (b builder) Scheme() string {
	return "easycar"
}
