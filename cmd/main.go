package main

import (
	"context"
	"log"
	"os"

	"github.com/wuqinqiang/easycar/services/coordinator"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Commands: []cli.Command{
			EasyCarCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

var EasyCarCommand = cli.Command{
	// todo add args
	Name: "easycar",
	Action: func(ctx *cli.Context) error {
		service, err := coordinator.New(coordinator.WithPort(8089))
		if err != nil {
			panic(err)
		}
		if err = service.Start(context.Background()); err != nil {
			panic(err)
		}
		return nil
	},
}
