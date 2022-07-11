package entity

import (
	"github.com/wuqinqiang/easycar/core/consts"
)

type Global struct {
	gId          string             // global id
	state        consts.GlobalState // global state
	endTime      int64              // end time for the transaction
	NextCronTime int64              // next cron time
}

func NewGlobal(gId string) *Global {
	return &Global{
		gId: gId,
	}
}

func (g *Global) IsEmpty() bool {
	return g.gId == ""
}

func (g *Global) IsCommitted() bool {
	return g.state == consts.Submitted
}

func (g *Global) IsCommitting() bool {
	return g.state == consts.Submitting
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

func (g *Global) GetEndTime() int64 {
	return g.endTime
}

func (g *Global) CanSubmit() bool {
	return g.state == consts.Begin
}

func (g *Global) GetBranches() []string {
	return []string{}
}
