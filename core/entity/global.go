package entity

import (
	"github.com/wuqinqiang/easycar/core/consts"
)

type Global struct {
	GID          string             `gorm:"column:g_id;type:varchar(255);not null" bson:"g_id"`                      // global id
	State        consts.GlobalState `gorm:"column:state;type:varchar(255);not null;default:init" bson:"state"`       // global State
	EndTime      int64              `gorm:"column:end_time;type:int;not null;default:0" bson:"end_time"`             // end time for the transaction
	NextCronTime int64              `gorm:"column:next_cron_time;type:int;not null;default:0" bson:"next_cron_time"` // next cron time
	CreateTime   int64              `gorm:"create_time;autoCreateTime" json:"create_time" bson:"create_time"`        // create time
	UpdateTime   int64              `gorm:"update_time;autoCreateTime" json:"update_time" bson:"update_time"`        // last update time
}

func (g Global) TableName() string {
	return "global"
}

func NewGlobal(gId string) *Global {
	return &Global{
		GID: gId,
	}
}

func (g *Global) Phase2Failed() bool {
	return g.State == consts.Phase1Failed
}

func (g *Global) Phase1Success() bool {
	return g.State == consts.Phase1Success
}

func (g *Global) Phase1Failed() bool {
	return g.State == consts.Phase1Failed
}

func (g *Global) Phase2Retrying() bool {
	return g.State == consts.Phase1Retrying
}
func (g *Global) Phase1Retrying() bool {
	return g.State == consts.Phase1Retrying
}

func (g *Global) Init() bool {
	return g.State == consts.Init
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

func (g *Global) AllowRegister() bool {
	return g.Init()
}
