package core

import "context"

type Server interface {
	Run(ctx context.Context) error
	Stop(ctx context.Context) error
}
