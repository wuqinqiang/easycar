package entity

import (
	"github.com/wuqinqiang/easycar/core/consts"
)

type Global struct {
	GID          string             `gorm:"column:g_id;type:varchar(255);not null"`                // global id
	State        consts.GlobalState `gorm:"column:state;type:varchar(255);not null;default:ready"` // global State
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

func (g *Global) CanPhase2() bool {
	return g.IsPhase2Failed()
}

func (g *Global) IsPhase2Failed() bool {
	return g.State == consts.Phase1Failed
}

func (g *Global) IsPhase1Success() bool {
	return g.State == consts.Phase1Success
}

func (g *Global) IsPhase2Retrying() bool {
	return g.State == consts.Phase1Retrying
}
func (g *Global) IsPhase1Retrying() bool {
	return g.State == consts.Phase1Retrying
}

func (g *Global) IsReady() bool {
	return g.State == consts.Ready
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

func (g *Global) AllowSubmit() bool {
	return g.IsReady()
}

func (g *Global) AllowRegister() bool {
	return g.IsReady()
}

func (g *Global) GetBranches() []string {
	return []string{}
}
