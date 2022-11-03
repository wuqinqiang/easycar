package etcdx

import (
	"context"
	"fmt"
	"strconv"

	"github.com/wuqinqiang/easycar/logging"

	"github.com/wuqinqiang/easycar/tools"

	"github.com/wuqinqiang/easycar/core/registry"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Registry struct {
	opts   Options
	client *clientv3.Client
}

func NewRegistry(conf Conf, fns ...Option) (*Registry, error) {
	opts := newDefault()

	for _, fn := range fns {
		fn(&opts)
	}

	r := &Registry{opts: opts}

	var (
		err error
	)
	if r.client, err = clientv3.New(clientv3.Config{
		Endpoints: conf.Hosts,
		Username:  conf.User,
		Password:  conf.Pass,
	}); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Registry) Register(ctx context.Context, instance *registry.EasyCarInstance) error {
	grant, err := r.client.Grant(ctx, int64(r.opts.ttl.Seconds()))
	if err != nil {
		return err
	}

	instance.Id = strconv.FormatInt(int64(grant.ID), 10)

	_, err = r.client.Put(ctx, instance.Key(), instance.Marshal(), clientv3.WithLease(grant.ID))
	if err != nil {
		return err
	}
	tools.GoSafe(func() {
		r.keepalive(ctx, grant.ID)
	})
	return nil
}

func (r *Registry) DeRegister(ctx context.Context, instance *registry.EasyCarInstance) error {
	_, err := r.client.Delete(ctx, instance.Key())
	return err
}

func (r *Registry) keepalive(ctx context.Context, id clientv3.LeaseID) {
	resp, err := r.client.KeepAlive(ctx, id)
	if err != nil {
		logging.Errorf(fmt.Sprintf("[keepalive] err:%v", err))
		return
	}

	for {
		select {
		case _, ok := <-resp:
			if !ok {
				if ctx.Err() != nil {
					logging.Errorf(fmt.Sprintf("[keepalive] resp err:%v", err))
					return
				}
			}
		case <-r.opts.ctx.Done():
			return
		}
	}

}

func (r *Registry) Watch(ctx context.Context, key string) (registry.Watcher, error) {
	w, err := NewWatcher(ctx, r.client, key)
	return w, err
}
