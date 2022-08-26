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
	Port: 8089,
}
