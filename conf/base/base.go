package base

import "github.com/wuqinqiang/easycar/conf"

var DefaultConf = conf.Settings{
	DB: conf.DB{
		Driver: "mysql",
		Mysql: conf.MysqlSettings{
			DbURL:        "easycar:easycar@tcp(127.0.0.1:3306)/easycar?charset=utf8&parseTime=True&loc=Local",
			MaxLifetime:  20,
			MaxIdleConns: 10,
			MaxOpenConns: 20,
		},
	},
	GRPCPort: 8089,
	HTTPPort: 8084,
}
