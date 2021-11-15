package handler

import (
	"github.com/wuqinqiang/easycar/internal/service"
)

type EasyCarHttpHandler struct {
	rm service.TMInterface
}

func NewEasyCarHttpHandler(rm service.TMInterface) EasyCarHttpHandler {
	return EasyCarHttpHandler{rm: rm}
}

func (http *EasyCarHttpHandler) run() {
	http.run()
}
