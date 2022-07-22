package main

import (
	"context"
	"log"
	"os"

	"github.com/wuqinqiang/easycar/conf"

	"github.com/wuqinqiang/easycar/services/coordinator"
)

func main() {
	c, err := conf.NewConf(os.Getenv("conf_mode"))
	if err != nil {
		log.Fatal(err)
	}
	// init conf
	easycar, err := c.Load()
	if err != nil {
		log.Fatal(err)
	}
	easycar.DB.Mysql.Init()
	service, err := coordinator.New(c, coordinator.WithPort(8089))
	if err != nil {
		panic(err)
	}
	if err = service.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}
