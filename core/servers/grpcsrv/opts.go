package grpcsrv

import (
	"crypto/tls"
	"time"

	"google.golang.org/grpc"
)

type Opt func(g *GrpcSrv)

func WithGrpcOpts(opts ...grpc.ServerOption) Opt {
	return func(g *GrpcSrv) {
		g.grpcOpts = append(g.grpcOpts, opts...)
	}
}

func WithTimeOut(t time.Duration) Opt {
	return func(g *GrpcSrv) {
		g.timeout = t
	}
}

func WithTls(tls *tls.Config) Opt {
	return func(g *GrpcSrv) {
		g.tls = tls
	}
}
