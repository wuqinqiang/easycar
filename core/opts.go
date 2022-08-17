package core

import (
	"time"

	"google.golang.org/grpc"
)

var defaultOpts = opts{
	timeout: 3 * time.Second,
}

type (
	OptsFn func(opts *opts)
)

type opts struct {
	port     int
	timeout  time.Duration
	grpcOpts []grpc.ServerOption
}

func WithPort(port int) OptsFn {
	return func(opts *opts) {
		opts.port = port
	}
}
