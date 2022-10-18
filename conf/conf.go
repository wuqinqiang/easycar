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

	Settings struct {
		DB                  DB               `yaml:"db"`
		GRPCPort            int              `yaml:"grpcPort"`
		HTTPPort            int              `yaml:"httpPort"`
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
