package base

import "github.com/wuqinqiang/easycar/conf"

var DefaultConf = conf.EasyCar{
	DB: conf.DB{
		Driver: "mysql",
		Mysql: conf.Mysql{
			DbURL:        "",
			MaxLifetime:  20,
			MaxIdleConns: 10,
			MaxOpenConns: 20,
		},
	},
	GRPCPort: 8089,
	HTTPPort: 8084,
}
