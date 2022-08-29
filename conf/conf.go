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

	EasyCar struct {
		DB       DB  `yaml:"db"`
		GRPCPort int `yaml:"grpcPort"`
		HTTPPort int `yaml:"httpPort"`
	}
)

type Conf interface {
	Load() (*EasyCar, error)
}
