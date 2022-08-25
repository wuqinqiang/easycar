package main

import (
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
	core, err := core.New(core.WithPort(8089), core.WithConf(easyCarConf))
	if err != nil {
		log.Fatal(err)
	}
	if err := core.Run(); err != nil {
		log.Fatal(err)
	}

	// everything is over
}

func flagConf() conf.Conf {
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
