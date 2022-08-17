package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/wuqinqiang/easycar/core"

	"github.com/wuqinqiang/easycar/conf/file"

	"github.com/wuqinqiang/easycar/conf"
)

func main() {
	f := flagConf()
	// init conf
	easyCarConf, err := f.Load()
	if err != nil {
		log.Fatal(err)
	}
	easyCarConf.DB.Mysql.Init()
	service, err := core.New(f, core.WithPort(8089))
	if err != nil {
		panic(err)
	}
	if err = service.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func flagConf() conf.Conf {
	var (
		c conf.Conf
	)
	confMod := flag.String("mod", os.Getenv("CONF_MOD"), "configuration module")
	switch conf.Mode(*confMod) {
	case conf.File:
		filePath := flag.String("f", GetEnvOrDefault("FILE_PATH", "conf.yml"), "configuration file")
		c = file.NewFile(*filePath)
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
