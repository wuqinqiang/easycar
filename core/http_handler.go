package core

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wuqinqiang/easycar/core/entity"
	"github.com/wuqinqiang/easycar/pkg/common"
)

type EasyCarHttpHandler struct {
	tm TMInterface
}

func NewEasyCarHttpHandler() EasyCarHttpHandler {
	return EasyCarHttpHandler{}
}

func (e *EasyCarHttpHandler) Begin(c *gin.Context) {
	var req common.GlobalData
	if err := BindJSONData(c, &req); err != nil {
		// todo
		return
	}
	// todo ctx in all request

	var (
		param entity.Global
	)
	param.SetProtocol(req.GetProtocol())
	param.SetTransactionName(req.GetTransactionName())

	begin, err := e.tm.Begin(context.TODO(), &param)
	if err != nil {
		return
	}
	fmt.Println(begin)
	// todo common resp
}

func BindJSONData(c *gin.Context, res interface{}) error {
	if err := c.BindJSON(res); err != nil {
		return err
	}
	return nil
}
