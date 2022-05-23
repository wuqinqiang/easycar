package model

import (
	"github.com/wuqinqiang/easycar/core/consts"
)

type Global struct {
	gId     string
	state   consts.GlobalState
	EndTime int64
}

func NewGlobal(gId string) *Global {
	return &Global{
		gId: gId,
	}
}

func (g *Global) SetGId(gId string) {
	g.gId = gId
}

func (g *Global) GetGId() string {
	return g.gId
}
func (g *Global) SetState(state consts.GlobalState) {
	g.state = state
}

func (g *Global) GetState() consts.GlobalState {
	return g.state
}

func (g *Global) CanSubmit() bool {
	return g.state == consts.Begin
}
