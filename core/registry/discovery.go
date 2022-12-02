package registry

import (
	"context"
)

type Discovery interface {
	Watch(ctx context.Context, serviceName string) (Watcher, error)
}

type Watcher interface {
	GetInstances() ([]*Instance, error)
	Next() ([]*Instance, error)
	Stop() error
}
