package core

import (
	"time"

	"google.golang.org/grpc"
)

var defaultOpts = opts{
	timeout:     3 * time.Second,
	grpcTimeout: 5 * time.Second,
}

type (
	OptsFn func(opts *opts)
)

type opts struct {
	grpcPort    int
	httpPort    int
	grpcTimeout time.Duration
	timeout     time.Duration
	grpcOpts    []grpc.ServerOption
}

func WithGrpcPort(port int) OptsFn {
	return func(opts *opts) {
		opts.grpcPort = port
	}
}

func WithHttpPort(port int) OptsFn {
	return func(opts *opts) {
		opts.httpPort = port
	}
}

func WithGrpcOpts(grpcOpts ...grpc.ServerOption) OptsFn {
	return func(opts *opts) {
		opts.grpcOpts = grpcOpts
	}
}
