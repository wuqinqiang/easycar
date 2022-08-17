package core

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/wuqinqiang/easycar/conf"
)

type Service struct {
	conf conf.Conf
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
	return
}

func (s *Service) Start(ctx context.Context) error {
	return nil
}

func (s *Service) Stop() (err error) {
	s.once.Do(func() {
		err = s.lis.Close()
	})
	return nil
}
