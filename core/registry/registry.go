package registry

import (
	"context"
	"encoding/json"
	"fmt"
)

const (
	ServerName string = "easycar"
	Version    string = "v1"
)

type EasyCarInstance struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	// examples:
	// http://127.0.0.1:8080
	// grpc://127.0.0.1:8085
	Nodes []string `json:"node"`
}

func (instance *EasyCarInstance) String() string {
	return fmt.Sprintf("%+v", *instance)
}

func NewEasyCarInstance() *EasyCarInstance {
	return &EasyCarInstance{
		Name:    ServerName,
		Version: Version,
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
	return fmt.Sprintf("/%s/%s", instance.Name, instance.Id)
}

type Registry interface {
	Register(ctx context.Context, instance *EasyCarInstance) error
	DeRegister(ctx context.Context, instance *EasyCarInstance) error
}
