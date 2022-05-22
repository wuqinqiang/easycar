package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type App struct {
	engine *gin.Engine
	// todo grpc handler
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
