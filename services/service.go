package services

import "context"

type Service interface {
	Start(ctx context.Context) error
	Stop() error
}
