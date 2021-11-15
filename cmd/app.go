package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wuqinqiang/easycar/internal/handler"
	"time"
)

type App struct {
	engine *gin.Engine
	Cof    *Config
	http   *handler.EasyCarHttpHandler
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
	app.engine.POST("/easycar/begin", func(context *gin.Context) {
	})
}

func (app *App) Run() {
	go app.engine.Run()
}
