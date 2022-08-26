package envx

import (
	"os"
	"strconv"

	"github.com/wuqinqiang/easycar/conf/base"

	"github.com/wuqinqiang/easycar/conf"
)

type Env struct {
}

func (env *Env) Load() (*conf.EasyCar, error) {
	defaultConf := base.DefaultConf

	driver := os.Getenv("DB_DRIVER")
	if driver == "" {
		defaultConf.DB.Driver = driver
	}

	convertFn := func(item string) int {
		if item == "" {
			return 0
		}
		res, err := strconv.Atoi(item)
		if err != nil {
			return 0
		}
		return res
	}

	servicePort := convertFn("PORT")
	if servicePort > 0 {
		defaultConf.GRPCPort = servicePort
	}

	lifeTime := convertFn(os.Getenv("MYSQL_MAX_LIFE_TIME"))
	if lifeTime > 0 {
		defaultConf.DB.Mysql.MaxLifetime = lifeTime
	}

	maxIdleConn := convertFn(os.Getenv("MYSQL_MAX_IDLE_CONN"))
	if maxIdleConn > 0 {
		defaultConf.DB.Mysql.MaxIdleConns = maxIdleConn
	}

	maxOpenConn := convertFn(os.Getenv("MYSQL_MAX_OPEN_CONN"))
	if maxOpenConn > 0 {
		defaultConf.DB.Mysql.MaxOpenConns = maxOpenConn
	}
	return &defaultConf, nil
}
