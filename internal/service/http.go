package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

var DefaultEasyCarHttpService *EasyCarHttpService

func init() {
	e := gin.Default()
	DefaultEasyCarHttpService = &EasyCarHttpService{e}
	e.Use(func(c *gin.Context) {
		nowTime := time.Now()
		c.Next()
		fmt.Printf("request use :%v\n", time.Since(nowTime).Seconds())
	})
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{"msg": "pong"})
	})
}

type EasyCarHttpService struct {
	*gin.Engine
}

func (http *EasyCarHttpService) start() {
}

func (http *EasyCarHttpService) RegisterRouter() {
	http.POST("/easycar/begin")
}
