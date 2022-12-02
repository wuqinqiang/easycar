package consulx

import (
	"context"
	"strings"
	"sync/atomic"
	"time"

	"github.com/wuqinqiang/easycar/logging"

	"github.com/hashicorp/consul/api"

	"github.com/wuqinqiang/easycar/core/registry"
)

type watcher struct {
	client      *client
	ctx         context.Context
	serviceName string
	services    *atomic.Value
	ticker      *time.Ticker
	ch          chan struct{}
	waitTime    time.Duration
	waitIndex   uint64
}

func newWatcher(ctx context.Context, cli *client, serviceName string) *watcher {
	w := &watcher{
		client:      cli,
		ctx:         ctx,
		serviceName: serviceName,
		services:    &atomic.Value{},
		ticker:      time.NewTicker(4 * time.Second),
		ch:          make(chan struct{}, 1),
	}
	go w.loop()
	return w
}

func (w *watcher) loop() {
	defer w.ticker.Stop()
	for {
		select {
		case <-w.ticker.C:
			servers, err := w.GetInstances()
			if err != nil {
				time.Sleep(time.Second)
				continue
			}
			w.services.Store(servers)
			w.ch <- struct{}{}
		case <-w.ctx.Done():
			return
		}
	}
}

func (w *watcher) GetInstances() (services []*registry.Instance, err error) {
	query := &api.QueryOptions{
		WaitTime:  w.waitTime,
		WaitIndex: w.waitIndex,
	}

	var (
		serviceEntry []*api.ServiceEntry
		queryMeta    *api.QueryMeta
	)
	serviceEntry, queryMeta, err = w.client.service(w.serviceName, query)
	if err != nil {
		logging.Errorf("[consul] watcher err:%v", err)
		return
	}

	if queryMeta.LastIndex == w.waitIndex || len(serviceEntry) == 0 {
		return
	}
	services, err = w.convertFn()(serviceEntry)
	atomic.SwapUint64(&w.waitIndex, queryMeta.LastIndex)
	return
}

func (w *watcher) convertFn() func(entryList []*api.ServiceEntry) (
	services []*registry.Instance, err error) {
	return func(entryList []*api.ServiceEntry) (services []*registry.Instance, err error) {
		for _, entry := range entryList {
			var (
				version string
				nodes   []string
			)
			if len(entry.Service.Tags) > 0 {
				versionTag := strings.Split(entry.Service.Tags[0], "=")
				if len(versionTag) > 0 {
					version = versionTag[1]
				}
			}

			for _, tag := range entry.Service.TaggedAddresses {
				nodes = append(nodes, tag.Address)
			}

			services = append(services, &registry.Instance{
				Id:      entry.Service.ID,
				Name:    entry.Service.Service,
				Version: version,
				Nodes:   nodes,
			})
		}
		return
	}
}

func (w *watcher) Next() ([]*registry.Instance, error) {
	select {
	case <-w.ch:
	case <-w.ctx.Done():
		return nil, w.ctx.Err()
	}
	list, _ := w.services.Load().([]*registry.Instance)
	return list, nil
}

func (w *watcher) Stop() error {
	return nil
}
