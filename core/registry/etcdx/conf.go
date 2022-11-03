package etcdx

import (
	"context"
	"time"
)

// Conf is the config item with the given key on etcd.
type Conf struct {
	Hosts []string `yaml:"hosts"`
	User  string   `yaml:"user"`
	Pass  string   `yaml:"pass"`
	//CertFile           string `json:",optional"`
	//CertKeyFile        string `json:",optional=CertFile"`
	//CACertFile         string `json:",optional=CertFile"`
	InsecureSkipVerify bool `json:""`
}

func (c *Conf) IsEmpty() bool {
	return len(c.Hosts) == 0
}

type Option func(options *Options)

type Options struct {
	ttl time.Duration
	ctx context.Context
}

func newDefault() Options {
	return Options{ttl: 10 * time.Second, ctx: context.Background()}
}
