package grpc

import (
	"net"

	"google.golang.org/grpc"
)

type Service struct {
	*grpc.Server
	opts opts
	lis  net.Listener
}

func New(fns ...OptsFn) *Service {
	ots := defaultOpts
	for _, fn := range fns {
		fn(&ots)
	}
	s := &Service{
		opts: ots,
	}
	return s
}

func (g Service) Start() error {
	//TODO implement me
	panic("implement me")
}

func (g Service) Stop() error {
	//TODO implement me
	panic("implement me")
}
