package consulx

import "github.com/hashicorp/consul/api"

type Conf struct {
	api.Config
	Default bool
}

func (c *Conf) Empty() bool {
	return !c.Default && c.Address == ""
}

func (c *Conf) Conf() *api.Config {
	if c.Default {
		return api.DefaultConfig()
	}
	return &c.Config
}
