package main

import (
	"context"
	"crypto/tls"
	"flag"
	"log"
	"os"

	"github.com/wuqinqiang/easycar/tools"

	"github.com/wuqinqiang/easycar/logging"

	"github.com/wuqinqiang/easycar/core/coordinator/executor"

	"github.com/wuqinqiang/easycar/core/coordinator"

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
	// load conf settings from (file|etcd|env......)
	settings, err := c.Load()
	if err != nil {
		log.Fatal(err)
	}

	// init config:db conf.....
	settings.Init()

	// Create a Coordinator,The core logic is here.
	newCoordinator := coordinator.NewCoordinator(dao.GetTransaction(),
		executor.NewExecutor(), settings.AutomaticExecution2)

	// New grpc server
	var grpcOpts []grpcsrv.Option
	if settings.Grpc.Tls() {
		certificate, err := tls.LoadX509KeyPair(settings.Grpc.CertFile, settings.Grpc.KeyFile)
		if err != nil {
			log.Fatal(err)
		}
		grpcOpts = append(grpcOpts, grpcsrv.WithTls(&tls.Config{Certificates: []tls.Certificate{certificate}}))
	}
	grpcSrv, err := grpcsrv.New(tools.FigureOutListen(settings.Grpc.ListenOn), newCoordinator, grpcOpts...)
	if err != nil {
		log.Fatal(err)
	}

	var (
		opts []core.Option
	)

	opts = append(opts, core.WithServers(grpcSrv))

	if settings.Grpc.IsOpenGateway() {
		httpProxySrv := httpsrv.New(tools.FigureOutListen(settings.Http.ListenOn),
			grpcSrv.HttpHandler(settings.Grpc.Gateway.CertFile, settings.Grpc.Gateway.ServerName))
		opts = append(opts, core.WithServers(httpProxySrv))
	}

	if !settings.IsRegistryEmpty() {
		registry, err := settings.GetRegistry()
		if err != nil {
			log.Fatal(err)
		}
		opts = append(opts, core.WithRegistry(registry))
	}

	core := core.New(opts...)
	if err != nil {
		log.Fatal(err)
	}
	if err := core.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
	logging.Infof("easycar server is stopped")
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
