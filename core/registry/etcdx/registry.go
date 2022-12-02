package etcdx

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/wuqinqiang/easycar/logging"

	"github.com/wuqinqiang/easycar/tools"

	"github.com/wuqinqiang/easycar/core/registry"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Registry struct {
	opts   Options
	client *clientv3.Client
}

func New(conf Conf, fns ...Option) (*Registry, error) {
	opts := newDefault()
	for _, fn := range fns {
		fn(&opts)
	}
	r := &Registry{opts: opts}

	var err error

	etcdConf := clientv3.Config{
		Endpoints:   conf.Hosts,
		Username:    conf.User,
		Password:    conf.Pass,
		DialTimeout: 5 * time.Second,
	}
	if r.client, err = clientv3.New(etcdConf); err != nil {
		return nil, err
	}
	// checkout status
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = r.client.Status(ctx, etcdConf.Endpoints[0])
	return r, err
}

func (r *Registry) Register(ctx context.Context, instance *registry.Instance) error {
	grant, err := r.client.Grant(ctx, int64(r.opts.ttl.Seconds()))
	if err != nil {
		return err
	}

	if instance.Id == "" {
		instance.Id = strconv.FormatInt(int64(grant.ID), 10)
	}

	_, err = r.client.Put(ctx, instance.InstanceName(), instance.Marshal(), clientv3.WithLease(grant.ID))
	if err != nil {
		return err
	}
	tools.GoSafe(func() {
		r.keepalive(ctx, grant.ID)
	})
	return nil
}

func (r *Registry) DeRegister(ctx context.Context, instance *registry.Instance) error {
	_, err := r.client.Delete(ctx, instance.InstanceName())
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

func (r *Registry) Watch(ctx context.Context, serviceName string) (registry.Watcher, error) {
	w, err := NewWatcher(ctx, r.client, serviceName)
	return w, err
}
