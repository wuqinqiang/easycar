package registry

import (
	"context"
	"encoding/json"
	"fmt"
)

type EasyCarInstance struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	// examples:
	// http://127.0.0.1:8080
	// grpc://127.0.0.1:8085
	Node []string `json:"node"`
}

func NewEasyCarInstance() *EasyCarInstance {
	return &EasyCarInstance{
		Name:    "/easycar",
		Version: "v1",
	}
}

func Unmarshal(val []byte) (*EasyCarInstance, error) {
	var (
		instance EasyCarInstance
	)
	err := json.Unmarshal(val, &instance)
	return &instance, err
}

func (instance *EasyCarInstance) Marshal() string {
	val, _ := json.Marshal(instance)
	return string(val)
}

func (instance *EasyCarInstance) Key() string {
	return fmt.Sprintf("%s/%s", instance.Name, instance.Id)
}

type Registry interface {
	Register(ctx context.Context, instance *EasyCarInstance) error
	DeRegister(ctx context.Context, instance *EasyCarInstance) error
}

type Discovery interface {
	Watch(ctx context.Context, key string) (Watcher, error)
}

type Watcher interface {
	Next() ([]*EasyCarInstance, error)
	Stop() error
}
