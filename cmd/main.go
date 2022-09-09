package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/wuqinqiang/easycar/core/coordinator/executor"

	"github.com/wuqinqiang/easycar/core/coordinator"

	"github.com/wuqinqiang/easycar/core/transport/common"

	"github.com/wuqinqiang/easycar/core/dao"

	"github.com/wuqinqiang/easycar/core/servers/runner"

	"github.com/wuqinqiang/easycar/core/servers/httpsrv"

	"github.com/wuqinqiang/easycar/core/servers/grpcsrv"

	"github.com/wuqinqiang/easycar/conf/envx"

	"github.com/wuqinqiang/easycar/core"

	"github.com/wuqinqiang/easycar/conf/file"

	"github.com/wuqinqiang/easycar/conf"
)

func main() {
	conf := getConf()
	// init conf
	settings, err := conf.Load()
	if err != nil {
		log.Fatal(err)
	}
	MustLoad(settings)
	e := executor.NewExecutor()
	coordinator := coordinator.NewCoordinator(dao.GetTransaction(), e, settings.AutomaticExecution2)
	grpcSrv, err := grpcsrv.New(settings.GRPCPort, coordinator)
	if err != nil {
		log.Fatal(err)
	}
	httpProxySrv := httpsrv.New(settings.HTTPPort, settings.GRPCPort)
	//runner
	runnerSrv, err := runner.NewRunner("@every 50m", func(ctx context.Context) {
		fmt.Println("hello world")
	})

	core := core.New(core.WithServers(grpcSrv, httpProxySrv, runnerSrv))
	if err != nil {
		log.Fatal(err)
	}
	if err := core.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
	// everything is over
}

func MustLoad(settings *conf.Settings) {

	settings.DB.Mysql.Init()

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
		// todo don't not be so rude!!
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
