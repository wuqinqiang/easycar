package conf

import "fmt"

type EasyCar struct {
	*Server
	*Mysql
}

func (conf *EasyCar) GetServer() (*Server, error) {
	if conf.Server == nil {
		return nil, fmt.Errorf("server is nil")
	}
	if conf.Server.Addr == "" {
		conf.Server.Addr = "8089"
	}
	return conf.Server, nil
}

func (conf *EasyCar) GetMysql() (*Mysql, error) {
	if conf.Mysql == nil {
		return nil, fmt.Errorf("mysql is nil")
	}
	if conf.Mysql.DbURL == "" {
		return nil, fmt.Errorf("mysql.dburl is nil")
	}
	if conf.MaxLifetime == 0 {
		conf.MaxLifetime = 3600
	}
	if conf.Mysql.MaxIdleConns == 0 {
		conf.Mysql.MaxIdleConns = 20
	}

	if conf.Mysql.MaxOpenConns == 0 {
		conf.Mysql.MaxOpenConns = 10
	}
	return conf.Mysql, nil
}

// Server conf
type Server struct {
	Addr string `json:"addr" yaml:"addr"`
}

// Mysql conf
type Mysql struct {
	DbURL        string `json:"dbURL" yaml:"dbURL"`
	MaxLifetime  int64  `json:"maxLifetime" yaml:"maxLifetime"`
	MaxIdleConns int64  `json:"maxIdleConns" yaml:"maxIdleConns"`
	MaxOpenConns int64  `json:"maxOpenConns" yaml:"maxOpenConns"`
}

type Conf interface {
	Load() (Loader, error)
}

type Loader interface {
	GetServer() (*Server, error)
	//Mysql  todo replace,more db type
	GetMysql() (*Mysql, error)
}
