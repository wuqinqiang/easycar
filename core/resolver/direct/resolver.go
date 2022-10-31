package direct

import "google.golang.org/grpc/resolver"

type Resolver int

func (A Resolver) ResolveNow(options resolver.ResolveNowOptions) {}

func (A Resolver) Close() {}
