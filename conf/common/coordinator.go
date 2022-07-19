package common

import "fmt"

type EasyCar struct {
	Server *Server
	Mysql  *Mysql
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
	if conf.Mysql == nil {
		return nil, fmt.Errorf("Mysql is nil")
	}
	if conf.Mysql.DbURL == "" {
		return nil, fmt.Errorf("Mysql.dburl is nil")
	}
	if conf.Mysql.MaxLifetime == 0 {
		conf.Mysql.MaxLifetime = 3600
	}
	if conf.Mysql.MaxIdleConns == 0 {
		conf.Mysql.MaxIdleConns = 20
	}

	if conf.Mysql.MaxOpenConns == 0 {
		conf.Mysql.MaxOpenConns = 10
	}
	return conf.Mysql, nil
}
