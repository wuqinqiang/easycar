package registry

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

const (
	ServerName string = "easycar"
	Version    string = "v1"
)

type Instance struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	// examples:
	// http://127.0.0.1:8080
	// grpc://127.0.0.1:8085
	Nodes []string `json:"node"`
}

func (instance *Instance) String() string {
	return fmt.Sprintf("%+v", *instance)
}

func NewInstance() *Instance {
	return &Instance{
		Name:    ServerName,
		Version: Version,
		Id:      uuid.NewString(),
	}
}

func Unmarshal(val []byte) (*Instance, error) {
	var (
		instance Instance
	)
	err := json.Unmarshal(val, &instance)
	return &instance, err
}

func (instance *Instance) Marshal() string {
	val, _ := json.Marshal(instance)
	return string(val)
}

func (instance *Instance) InstanceName() string {
	return fmt.Sprintf("/%s/%s", instance.Name, instance.Id)
}

type Registry interface {
	Register(ctx context.Context, instance *Instance) error
	DeRegister(ctx context.Context, instance *Instance) error
}
