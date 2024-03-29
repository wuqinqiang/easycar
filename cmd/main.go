package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/wuqinqiang/easycar/core/notify"

	"github.com/wuqinqiang/easycar/core/servers/runner"

	"github.com/wuqinqiang/easycar/logging"

	"github.com/wuqinqiang/easycar/core/coordinator/executor"

	"github.com/wuqinqiang/easycar/core/coordinator"

	"github.com/wuqinqiang/easycar/core/dao"

	"github.com/wuqinqiang/easycar/core/servers/httpsrv"

	"github.com/wuqinqiang/easycar/core/servers/grpcsrv"

	"github.com/wuqinqiang/easycar/core"

	"github.com/wuqinqiang/easycar/conf/file"
)

var (
	filePath = flag.String("f", GetEnvOrDefault("FILE_PATH", "/conf.yml"), "configuration file")
)

func main() {
	flag.Parse()
	c := file.NewFile(*filePath)

	// load conf settings from (file|etcd|env......)
	settings, err := c.Load()
	if err != nil {
		log.Fatal(err)
	}

	// init config:db conf.....
	settings.Init()

	// Create a Coordinator,The core logic is here.
	dao := dao.GetTransaction()
	n := notify.New(settings.Notify.Senders())
	newCoordinator := coordinator.NewCoordinator(dao,
		executor.NewExecutor(), n, settings.AutomaticExecution2)

	var (
		servers []core.Server
	)
	// create grpc server
	grpcSrv, err := grpcsrv.New(settings.Grpc, newCoordinator)
	if err != nil {
		log.Fatal(err)
	}
	// cron server
	cronServer := runner.New(newCoordinator, dao,
		runner.WithMaxTimes(settings.Cron.MaxTimes), runner.WithTimeInterval(settings.Cron.TimeInterval))
	servers = append(servers, grpcSrv, cronServer)

	// create http server if needed
	if settings.Grpc.OpenGateway() {
		httpProxySrv := httpsrv.New(settings.Http,
			grpcSrv.Handler(settings.Grpc.Gateway.CertFile, settings.Grpc.Gateway.ServerName))
		servers = append(servers, httpProxySrv)
	}

	var (
		opts []core.Option
	)
	opts = append(opts, core.WithServers(servers...))

	// If the registry is configured,
	//we need to register the service to the  registry center when the server start
	if settings.SetRegistry() {
		registry, err := settings.GetRegistry()
		if err != nil {
			log.Fatal(err)
		}
		opts = append(opts, core.WithRegistry(registry))
	}

	// servers start-up core
	newCore := core.New(opts...)
	if err != nil {
		log.Fatal(err)
	}
	if err := newCore.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
	logging.Infof("easycar server is stopped")
}

func GetEnvOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
