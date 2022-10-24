package conf

import (
	"fmt"

	"github.com/wuqinqiang/easycar/core/registry"

	"github.com/wuqinqiang/easycar/core/registry/etcdx"
)

type (
	Mode string
)

const (
	File Mode = "file"
	Etcd Mode = "etcd"
	Env  Mode = "env"
	//Add more conf schema here
)

type (
	DB struct {
		Driver  string        `yaml:"driver"`
		Mysql   MysqlSettings `yaml:"mysql"`
		Mongodb MongoSetting  `yaml:"mongodb"`
	}

	Server struct {
		Http struct {
			ListenOn string `yaml:"listenOn"`
		} `yaml:"http"`
		Grpc Grpc `yaml:"grpc"`
	}

	Grpc struct {
		ListenOn string  `yaml:"listenOn"`
		KeyFile  string  `yaml:"keyFile"`
		CertFile string  `yaml:"certFile"`
		Gateway  Gateway `yaml:"gateway"`
	}
	Gateway struct {
		IsOpen     bool   `yaml:"isOpen"`
		CertFile   string `yaml:"certFile"`
		ServerName string `yaml:"serverName"`
	}

	Settings struct {
		Server              `json:"server"`
		DB                  DB               `yaml:"db"`
		Timeout             int64            `yaml:"timeout"`
		AutomaticExecution2 bool             `yaml:"automaticExecution2"`
		Tracing             Tracing          `yaml:"tracing"`
		Registry            RegistrySettings `yaml:"registry"`
	}

	RegistrySettings struct {
		Etcd etcdx.Conf `yaml:"etcd"`
	}

	Tracing struct {
		JaegerUri string `yaml:"jaegerUrl"`
	}
)

type Conf interface {
	Load() (*Settings, error)
}

func (db *DB) Init() {
	switch db.Driver {
	case "mysql":
		db.Mysql.Init()
	case "mongodb":
		db.Mongodb.Init()
	default:
		panic(fmt.Errorf("no support %s database", db.Driver))
	}
}

func (s *Settings) IsRegistryEmpty() bool {
	return s.Registry.Etcd.IsEmpty()
}
func (s *Settings) GetRegistry() (registry.Registry, error) {
	if !s.Registry.Etcd.IsEmpty() {
		return etcdx.NewRegistry(s.Registry.Etcd)
	}
	return nil, nil
}

func (grpc *Grpc) IsSSL() bool {
	return grpc.KeyFile != "" && grpc.CertFile != ""
}

func (grpc *Grpc) IsOpenGateway() bool {
	return grpc.Gateway.IsOpen
}
