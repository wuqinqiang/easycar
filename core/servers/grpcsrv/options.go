package grpcsrv

import (
	"crypto/tls"
	"time"

	"google.golang.org/grpc"
)

type Option func(g *GrpcSrv)

func WithGrpcOpts(opts ...grpc.ServerOption) Option {
	return func(g *GrpcSrv) {
		g.grpcOpts = append(g.grpcOpts, opts...)
	}
}

func WithTimeOut(t time.Duration) Option {
	return func(g *GrpcSrv) {
		g.timeout = t
	}
}

func WithTls(tls *tls.Config) Option {
	return func(g *GrpcSrv) {
		g.tls = tls
	}
}
