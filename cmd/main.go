package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/wuqinqiang/easycar/core/registry/etcdx"

	"github.com/wuqinqiang/easycar/tracing"

	"github.com/wuqinqiang/easycar/logging"

	"github.com/wuqinqiang/easycar/core/coordinator/executor"

	"github.com/wuqinqiang/easycar/core/coordinator"

	"github.com/wuqinqiang/easycar/core/transport/common"

	"github.com/wuqinqiang/easycar/core/dao"

	"github.com/wuqinqiang/easycar/core/servers/httpsrv"

	"github.com/wuqinqiang/easycar/core/servers/grpcsrv"

	"github.com/wuqinqiang/easycar/conf/envx"

	"github.com/wuqinqiang/easycar/core"

	"github.com/wuqinqiang/easycar/conf/file"

	"github.com/wuqinqiang/easycar/conf"
)

func main() {
	c := getConf()
	// init conf
	settings, err := c.Load()
	if err != nil {
		log.Fatal(err)
	}
	MustLoad(settings)
	setCoordinator := coordinator.NewCoordinator(dao.GetTransaction(),
		executor.NewExecutor(), settings.AutomaticExecution2)
	grpcSrv, err := grpcsrv.New(settings.GRPCPort, setCoordinator)
	if err != nil {
		log.Fatal(err)
	}
	httpProxySrv := httpsrv.New(settings.HTTPPort, settings.GRPCPort)

	etcdRegistry, err := etcdx.NewRegistry(etcdx.Conf{
		Hosts:              []string{"127.0.0.1:2379"},
		InsecureSkipVerify: false,
	})
	if err != nil {
		log.Fatal(err)
	}

	core := core.New(core.WithServers(grpcSrv, httpProxySrv), core.WithRegistry(etcdRegistry))
	if err != nil {
		log.Fatal(err)
	}
	if err := core.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
	logging.Infof("easycar server is stopped")
}

func MustLoad(settings *conf.Settings) {
	settings.DB.Init()

	tracing.MustLoad(settings.Tracing.JaegerUri)

	if settings.Timeout > 0 {
		common.ReplaceTimeOut(time.Duration(settings.Timeout) * time.Second)
	}
}

func getConf() (c conf.Conf) {
	var (
		confMod  = flag.String("mod", os.Getenv("CONF_MOD"), "configuration module")
		filePath = flag.String("f", GetEnvOrDefault("FILE_PATH", "/conf.yml"), "configuration file")
	)
	flag.Parse()

	switch conf.Mode(*confMod) {
	case conf.File:
		c = file.NewFile(*filePath)
	case conf.Etcd:
	case conf.Env:
		return new(envx.Env)
	default:
		panic("conf mod not set")
	}
	return c
}

func GetEnvOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
