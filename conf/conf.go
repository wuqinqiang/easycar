package conf

import (
	"fmt"
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

type Conf interface {
	Load() (*EasyCar, error)
}

type DB struct {
	Mysql *Mysql `json:"mysql"`
}

type EasyCar struct {
	Dirver string  `ymal:"dirver"`
	DB     *DB     `yaml:"db"`
	Server *Server `yaml:"server"`
}

// Server conf
type Server struct {
	Addr string `json:"addr" yaml:"addr"`
}

func (conf *EasyCar) GetServer() (*Server, error) {
	if conf.Server == nil {
		return nil, fmt.Errorf("Server is nil")
	}
	if conf.Server.Addr == "" {
		conf.Server.Addr = "8089"
	}
	return conf.Server, nil
}

func (conf *EasyCar) GetMysql() (*Mysql, error) {
	if conf.DB.Mysql == nil {
		return nil, fmt.Errorf("Mysql is nil")
	}
	if conf.DB.Mysql.DbURL == "" {
		return nil, fmt.Errorf("Mysql.dburl is nil")
	}
	if conf.DB.Mysql.MaxLifetime == 0 {
		conf.DB.Mysql.MaxLifetime = 3600
	}
	if conf.DB.Mysql.MaxIdleConns == 0 {
		conf.DB.Mysql.MaxIdleConns = 20
	}

	if conf.DB.Mysql.MaxOpenConns == 0 {
		conf.DB.Mysql.MaxOpenConns = 10
	}
	return conf.DB.Mysql, nil
}
