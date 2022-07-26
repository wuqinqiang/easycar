package entity

import (
	"github.com/wuqinqiang/easycar/core/consts"
)

type Global struct {
	GID          string             `gorm:"column:g_id;type:varchar(255);not null"`                // global id
	State        consts.GlobalState `gorm:"column:state;type:varchar(255);not null;default:begin"` // global State
	EndTime      int64              `gorm:"column:end_time;type:int;not null;default:0"`           // end time for the transaction
	NextCronTime int64              `gorm:"column:next_cron_time;type:int;not null;default:0"`     // next cron time
}

func (g Global) TableName() string {
	return "global"
}

func NewGlobal(gId string) *Global {
	return &Global{
		GID: gId,
	}
}

func (g *Global) CanCommit() bool {
	return g.IsBegin() || g.IsRetrying()
}

func (g *Global) CanRollback() bool {
	return g.IsCommitFailed() || g.IsRollBackRetrying()
}

func (g *Global) IsCommitFailed() bool {
	return g.State == consts.GlobalCommitFailed
}

func (g *Global) IsRollBackRetrying() bool {
	return g.State == consts.GlobalRollBackRetrying
}

func (g *Global) IsBegin() bool {
	return g.State == consts.Begin
}

func (g *Global) IsRetrying() bool {
	return g.State == consts.GlobalCommitRetrying
}

func (g *Global) IsEmpty() bool {
	return g.GID == ""
}

func (g *Global) SetGId(gId string) {
	g.GID = gId
}

func (g *Global) GetGId() string {
	return g.GID
}
func (g *Global) SetState(state consts.GlobalState) {
	g.State = state
}

func (g *Global) GetState() consts.GlobalState {
	return g.State
}

func (g *Global) GetEndTime() int64 {
	return g.EndTime
}

func (g *Global) CanSubmit() bool {
	return g.State == consts.Begin
}

func (g *Global) GetBranches() []string {
	return []string{}
}
