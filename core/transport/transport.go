package transport

import (
	"context"
	"fmt"
	"sync"

	grpc2 "github.com/wuqinqiang/easycar/core/transport/grpc"
	http2 "github.com/wuqinqiang/easycar/core/transport/http"

	"github.com/wuqinqiang/easycar/logging"

	"errors"

	"github.com/wuqinqiang/easycar/core/transport/common"
)

var (
	NotFoundTransport = errors.New("not found transport")
)

type (
	Transporter interface {
		// GetType returns the type of the net transport
		GetType() common.Net
		Request(ctx context.Context, url string, req *common.Req) (*common.Resp, error)
		Close(ctx context.Context) error
	}
	Manager interface {
		GetTransporter(net common.Net) (Transporter, error)
		Close(ctx context.Context) error
	}
)

type manager struct {
	m sync.Map
}

func NewManager() *manager {
	manager := &manager{
		m: sync.Map{},
	}
	var (
		list []Transporter
	)
	list = append(list, http2.NewTransporter(), grpc2.NewTransporter())
	for _, transporter := range list {
		manager.m.Store(string(transporter.GetType()), transporter)
	}
	return manager
}

func (manager *manager) GetTransporter(net common.Net) (Transporter, error) {
	val, ok := manager.m.Load(string(net))
	if !ok {
		return nil, NotFoundTransport
	}
	return val.(Transporter), nil
}

func (manager *manager) Close(ctx context.Context) error {
	manager.m.Range(func(key, value any) bool {
		if err := value.(Transporter).Close(ctx); err != nil {
			logging.Infof(fmt.Sprintf("[Manager] stop err:%v", err), "net", key, "transporter", value)
		}
		logging.Infof("[Manager] close client connections", "net", key, "transporter", value)
		return true
	})
	return nil
}
