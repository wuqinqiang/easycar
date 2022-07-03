package grpc

import (
	"time"

	"google.golang.org/grpc"
)

type OptsFn func(opts *opts)

var defaultOpts = opts{
	timeout: 3 * time.Second,
}

type opts struct {
	timeout    time.Duration
	unaryInts  []grpc.UnaryServerInterceptor
	streamInts []grpc.StreamServerInterceptor
	grpcOpts   []grpc.ServerOption
}
