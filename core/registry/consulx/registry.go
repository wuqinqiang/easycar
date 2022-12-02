package consulx

import (
	"context"
	"time"

	"github.com/wuqinqiang/easycar/logging"

	"github.com/hashicorp/consul/api"
	"github.com/wuqinqiang/easycar/core/registry"
)

var _ registry.Discovery = (*Registry)(nil)

type Registry struct {
	client *client
	cancel func()
	ctx    context.Context
}

func New(client *api.Client) *Registry {
	r := &Registry{
		client: NewClient(client),
	}
	r.ctx, r.cancel = context.WithCancel(context.Background())
	return r
}

func (r *Registry) Register(ctx context.Context, instance *registry.Instance) error {
	checkId := "easycar:" + instance.Id
	if err := r.client.register(ctx, checkId, instance); err != nil {
		logging.Errorf("[consul] Register instance %+v, err:%v", instance, err)
		return err
	}
	go r.keepalive(checkId)
	return nil
}

func (r *Registry) DeRegister(ctx context.Context, instance *registry.Instance) error {
	return r.client.deRegister(ctx, instance)
}

func (r *Registry) Watch(ctx context.Context, serviceName string) (registry.Watcher, error) {
	return newWatcher(ctx, r.client, serviceName), nil
}

func (r *Registry) keepalive(checkId string) {
	time.Sleep(time.Second)
	err := r.client.updateTTL(context.Background(), checkId)
	if err != nil {
		logging.Errorf("[consul] UpdateTTL checkId:%v,err:%v", checkId, err)
	}

	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err := r.client.updateTTL(context.Background(), checkId)
			if err != nil {
				logging.Errorf("[consul] UpdateTTL checkId:%v,err:%v", checkId, err)
			}

		case <-r.ctx.Done():
			return
		}
	}
}
