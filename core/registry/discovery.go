package registry

import (
	"context"
)

type Discovery interface {
	Watch(ctx context.Context, key string) (Watcher, error)
}

type Watcher interface {
	Next() ([]*EasyCarInstance, error)
	Stop() error
}
