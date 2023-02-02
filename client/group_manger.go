package client

import "github.com/wuqinqiang/easycar/core/consts"

type Manger struct {
	groups []*Group
	// level current level
	level consts.Level
}

func NewManger() *Manger {
	return &Manger{
		// default level
		level: 1,
	}
}

func (m *Manger) AddGroup(incrLevel bool, groups ...*Group) *Manger {
	if incrLevel {
		m.level++
	}
	for _, group := range groups {
		group.SetLevel(m.level)
		m.groups = append(m.groups, group)
	}
	return m
}

func (m *Manger) Groups() []*Group {
	return m.groups
}
