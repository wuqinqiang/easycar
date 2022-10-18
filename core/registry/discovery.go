package registry

import (
	"context"
	"fmt"
	"sync"
)

var (
	m      map[string]Discovery
	locker sync.RWMutex
)

func init() {
	m = make(map[string]Discovery)
}

func RegisterDiscovery(key string, discovery Discovery) {
	if key == "" || discovery == nil {
		return
	}
	locker.Lock()
	defer locker.Unlock()
	_, ok := m[key]
	if ok {
		return
	}
	m[key] = discovery
}

func GetDiscovery(key string) (Discovery, error) {
	locker.RLocker()
	defer locker.RUnlock()
	d, ok := m[key]
	if !ok {
		return nil, fmt.Errorf("no discovery")
	}
	return d, nil
}

type Discovery interface {
	Watch(ctx context.Context, key string) (Watcher, error)
}

type Watcher interface {
	Next() ([]*EasyCarInstance, error)
	Stop() error
}
