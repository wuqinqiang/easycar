package model

import (
	"strings"

	"github.com/wuqinqiang/easycar/core/consts"
)

type Global struct {
	gId      string
	state    consts.GlobalState
	protocol string
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

func (g *Global) SetProtocol(protocol string) {
	g.protocol = protocol
}

func (g *Global) GetProtocol() string {
	return g.protocol
}

func (g *Global) IsGrpc() bool {
	return strings.HasSuffix(g.protocol, "grpc")
}

func (g *Global) CanSubmit() bool {
	return g.state == consts.Begin
}
