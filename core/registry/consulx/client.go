package consulx

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/wuqinqiang/easycar/logging"

	"github.com/hashicorp/consul/api"
	"github.com/wuqinqiang/easycar/core/registry"
)

type client struct {
	cli *api.Client
}

func NewClient(cli *api.Client) *client {
	return &client{
		cli: cli,
	}
}

func (client *client) register(ctx context.Context, checkId string, instance *registry.Instance) error {
	taggedAddress := make(map[string]api.ServiceAddress)
	for _, node := range instance.Nodes {
		u, err := url.Parse(node)
		if err != nil {
			logging.Errorf("[consul] Registry Node:%v,err:%v", node, err)
			continue
		}
		port, _ := strconv.ParseUint(u.Port(), 10, 16) //nolint:errcheck
		taggedAddress[u.Scheme] = api.ServiceAddress{
			Address: node,
			Port:    int(port),
		}
	}

	reg := &api.AgentServiceRegistration{
		ID:                instance.Id,
		Name:              instance.Name,
		TaggedAddresses:   taggedAddress,
		Tags:              []string{fmt.Sprintf("version=%s", instance.Version)},
		EnableTagOverride: true,
	}

	reg.Checks = append(reg.Checks, &api.AgentServiceCheck{
		CheckID:                        checkId,
		TTL:                            "20s",
		Name:                           instance.InstanceName(),
		DeregisterCriticalServiceAfter: "",
	})
	q := new(api.ServiceRegisterOpts)
	return client.cli.Agent().ServiceRegisterOpts(reg, q.WithContext(ctx))
}

func (client *client) deRegister(ctx context.Context, instance *registry.Instance) error {
	q := new(api.QueryOptions)
	return client.cli.Agent().ServiceDeregisterOpts(instance.Id, q.WithContext(ctx))
}

func (client *client) updateTTL(_ context.Context, checkId string) error {
	return client.cli.Agent().UpdateTTL(checkId, "pass", "pass")
}

func (client *client) service(server string, q *api.QueryOptions) ([]*api.ServiceEntry, *api.QueryMeta, error) {
	return client.cli.Health().Service(server, "", true, q)
}
