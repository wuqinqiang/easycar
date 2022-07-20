package common

import "fmt"

type Db struct {
	Mysql *Mysql `json:"mysql"`
}

type EasyCar struct {
	Db     *Db `json:"db"`
	Server *Server
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
	if conf.Db.Mysql == nil {
		return nil, fmt.Errorf("Mysql is nil")
	}
	if conf.Db.Mysql.DbURL == "" {
		return nil, fmt.Errorf("Mysql.dburl is nil")
	}
	if conf.Db.Mysql.MaxLifetime == 0 {
		conf.Db.Mysql.MaxLifetime = 3600
	}
	if conf.Db.Mysql.MaxIdleConns == 0 {
		conf.Db.Mysql.MaxIdleConns = 20
	}

	if conf.Db.Mysql.MaxOpenConns == 0 {
		conf.Db.Mysql.MaxOpenConns = 10
	}
	return conf.Db.Mysql, nil
}
