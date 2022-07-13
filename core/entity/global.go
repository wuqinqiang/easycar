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

func (g *Global) CanCommit() bool {
	return g.IsBegin() || g.IsRetrying()
}

func (g *Global) CanRollback() bool {
	return g.IsCommitFailed() || g.IsRollBackRetrying()
}

func (g *Global) IsCommitFailed() bool {
	return g.state == consts.GlobalCommitFailed
}

func (g *Global) IsRollBackRetrying() bool {
	return g.state == consts.GlobalRollBackRetrying
}

func (g *Global) IsBegin() bool {
	return g.state == consts.Begin
}

func (g *Global) IsRetrying() bool {
	return g.state == consts.GlobalCommitRetrying
}

func (g *Global) IsEmpty() bool {
	return g.gId == ""
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
