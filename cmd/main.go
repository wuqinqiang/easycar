package main

import (
	"flag"
	"log"
	"os"

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
	settings.DB.Mysql.Init()
	core, err := core.New(core.WithHttpPort(settings.HTTPPort), core.WithGrpcPort(settings.GRPCPort))
	if err != nil {
		log.Fatal(err)
	}
	if err := core.Run(); err != nil {
		log.Fatal(err)
	}
	// everything is over
}

func getConf() conf.Conf {
	var (
		c conf.Conf
	)
	confMod := flag.String("mod", os.Getenv("CONF_MOD"), "configuration module")
	switch conf.Mode(*confMod) {
	case conf.File:
		filePath := flag.String("f", GetEnvOrDefault("FILE_PATH", "/conf.yml"), "configuration file")
		c = file.NewFile(*filePath)
	case conf.Etcd:
	case conf.Env:
		return new(envx.Env)
	default:
		// todo don't not be so rude!!
		panic("conf mod not set")
	}
	flag.Parse()
	return c
}

func GetEnvOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
