package main

import (
	"fmt"
	"time"

	"github.com/wuqinqiang/easycar/core"

	"github.com/gin-gonic/gin"
)

type App struct {
	engine *gin.Engine
	Cof    *Config
	http   *core.EasyCarHttpHandler
	// todo grpc handler
}

type Config struct {
	Server Server `json:"server"`
}

// Server server config
type Server struct {
	url  string
	port int64
}

func NewApp() *App {
	e := gin.Default()
	e.Use(func(c *gin.Context) {
		nowTime := time.Now()
		c.Next()
		fmt.Printf("request use :%v\n", time.Since(nowTime).Seconds())
	})
	// todo wire
	return &App{engine: e}
}

func (app *App) RegisterRouter() {
	handler := core.NewEasyCarHttpHandler()
	app.engine.POST("/easycar/begin", func(context *gin.Context) {
		handler.Begin(context)
	})
}

func (app *App) Run() {
	go app.engine.Run()
}
