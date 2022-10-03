package conf

import "fmt"

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
		DB                  DB      `yaml:"db"`
		GRPCPort            int     `yaml:"grpcPort"`
		HTTPPort            int     `yaml:"httpPort"`
		Timeout             int64   `yaml:"timeout"`
		AutomaticExecution2 bool    `yaml:"automaticExecution2"`
		Retry               Retry   `yaml:"retry"`
		Tracing             Tracing `yaml:"tracing"`
	}
	Retry struct {
		MaxDelay uint32 ` yaml:"maxDelay"`
		Retries  uint32 ` yaml:"retries"`
		Factor   uint32 `yaml:"factor"`
		Open     bool   `yaml:"open"`
	}

	Tracing struct {
		JaegerUri string `yaml:"jaegerUrl"`
	}
)

func (r *Retry) IsOpen() bool {
	return r.Open
}

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
