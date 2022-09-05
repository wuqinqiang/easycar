package conf

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
		Driver string `ymal:"driver"`
		Mysql  Mysql  `json:"mysql"`
	}

	Settings struct {
		DB                  DB    `yaml:"db"`
		GRPCPort            int   `yaml:"grpcPort"`
		HTTPPort            int   `yaml:"httpPort"`
		Timeout             int64 `yaml:"timeout"`
		AutomaticExecution2 bool  `yaml:"automaticExecution2"`
	}
)

type Conf interface {
	Load() (*Settings, error)
}
