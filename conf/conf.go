package conf

import (
	"fmt"
	"time"

	"github.com/wuqinqiang/easycar/core/servers/httpsrv"

	"github.com/wuqinqiang/easycar/core/transport/common"

	"github.com/wuqinqiang/easycar/tracing"

	"github.com/wuqinqiang/easycar/core/servers/grpcsrv"

	"github.com/wuqinqiang/easycar/core/dao/mongodb"

	gormx "github.com/wuqinqiang/easycar/core/dao/gorm"

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
	Settings struct {
		Server              `yaml:"server"`
		DB                  DB               `yaml:"db"`
		Timeout             int64            `yaml:"timeout"`
		AutomaticExecution2 bool             `yaml:"automaticExecution2"`
		Tracing             Tracing          `yaml:"tracing"`
		Registry            RegistrySettings `yaml:"registry"`
	}

	DB struct {
		Driver  string           `yaml:"driver"`
		Mysql   gormx.Settings   `yaml:"mysql"`
		Mongodb mongodb.Settings `yaml:"mongodb"`
	}

	Server struct {
		Http httpsrv.Http `yaml:"http"`
		Grpc grpcsrv.Grpc `yaml:"grpc"`
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

func (s *Settings) Init() {
	s.DB.Init()
	tracing.Init(s.Tracing.JaegerUri)

	if s.Http.ListenOn == "" {
		s.Http.ListenOn = "0.0.0.0:8085"
	}
	if s.Grpc.ListenOn == "" {
		s.Grpc.ListenOn = "0.0.0.0:8088"
	}
	if s.Timeout > 0 {
		common.ReplaceTimeout(time.Duration(s.Timeout) * time.Second)
	}
}

func (s *Settings) EmptyRegistry() bool {
	return s.Registry.Etcd.IsEmpty()
	// todo add more registry center
}
func (s *Settings) GetRegistry() (registry.Registry, error) {
	if !s.Registry.Etcd.IsEmpty() {
		return etcdx.NewRegistry(s.Registry.Etcd)
	}
	return nil, nil
}
