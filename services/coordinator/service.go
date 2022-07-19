package coordinator

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/wuqinqiang/easycar/conf"

	"github.com/wuqinqiang/easycar/core/dao"

	"github.com/wuqinqiang/easycar/core"

	"github.com/wuqinqiang/easycar/proto"

	"google.golang.org/grpc"
)

type Service struct {
	conf conf.Conf
	*grpc.Server
	opts opts
	lis  net.Listener
	once sync.Once
}

func New(conf conf.Conf, fns ...OptsFn) (s *Service, err error) {
	opts := defaultOpts
	for _, fn := range fns {
		fn(&opts)
	}
	s = &Service{
		conf: conf,
		opts: opts,
		once: sync.Once{},
	}
	if s.lis, err = net.Listen("tcp", fmt.Sprintf(":%d", opts.port)); err != nil {
		return
	}

	s.Server = grpc.NewServer(opts.grpcOpts...)
	c := core.NewCoordinator(dao.GetTransaction())
	proto.RegisterEasyCarServer(s.Server, NewCoordinator(c))
	return
}

func (s *Service) Start(ctx context.Context) error {
	return s.Server.Serve(s.lis)
}

func (s *Service) Stop() (err error) {
	s.once.Do(func() {
		err = s.lis.Close()
	})
	return nil
}
