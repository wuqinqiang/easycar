package coordinator

import (
	"time"

	"google.golang.org/grpc"
)

type OptsFn func(opts *opts)

var defaultOpts = opts{
	timeout: 3 * time.Second,
}

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
