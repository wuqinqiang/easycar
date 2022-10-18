package etcdx

import (
	"context"

	"github.com/wuqinqiang/easycar/core/registry"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type watcher struct {
	first     bool
	ctx       context.Context
	cancel    func()
	key       string
	client    *clientv3.Client
	watchChan clientv3.WatchChan
}

func NewWatcher(ctx context.Context, client *clientv3.Client, key string) (*watcher, error) {
	w := &watcher{
		first:  true,
		key:    key,
		client: client,
	}
	w.ctx, w.cancel = context.WithCancel(ctx)
	w.watchChan = w.client.Watch(ctx, key, clientv3.WithPrefix(), clientv3.WithRev(0))
	if err := w.client.Watcher.RequestProgress(ctx); err != nil {
		return nil, err
	}
	return w, nil
}

func (w *watcher) Next() ([]*registry.EasyCarInstance, error) {
	if !w.first {
		w.first = false
		return w.getInstances()
	}

	select {
	case <-w.ctx.Done():
		return nil, w.ctx.Err()
	case resp := <-w.watchChan:
		if resp.Err() != nil {
			return nil, resp.Err()
		}
		return w.getInstances()
	}
}

func (w *watcher) getInstances() ([]*registry.EasyCarInstance, error) {
	resp, err := w.client.KV.Get(w.ctx, w.key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	var (
		list []*registry.EasyCarInstance
	)
	for _, kv := range resp.Kvs {
		instance, err := registry.Unmarshal(kv.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, instance)
	}
	return list, nil
}

func (w *watcher) Stop() error {
	w.cancel()
	return nil
}
