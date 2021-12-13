package entity

import (
	"strings"

	"github.com/wuqinqiang/easycar/pkg/common"
)

type Global struct {
	gId             string
	transactionName common.TransactionName
	state           common.GlobalState
	protocol        string
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

func (g *Global) SetTransactionName(transactionName common.TransactionName) {
	g.transactionName = transactionName
}

func (g *Global) GetTransactionName() common.TransactionName {
	return g.transactionName
}

func (g *Global) SetState(state common.GlobalState) {
	g.state = state
}

func (g *Global) GetState() common.GlobalState {
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
	return g.state == common.Prepared
}

func (g *Global) CanRollBack() bool {
	return g.state == common.Prepared
}
