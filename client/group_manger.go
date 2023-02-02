package client

import (
	"sync/atomic"

	"github.com/wuqinqiang/easycar/core/consts"
)

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

// AddNextWaitGroups
// The group added this time will be executed after the completion of the call to the group added by the previous level call to AddNextWaitGroup
// example:m.AddGroups(group1,group2).AddNextWaitGroups(group3).AddGroups(group4),that means:
// first concurrent execution group1 and group2,then execution the group3,finally execution the group4
func (m *Manger) AddNextWaitGroups(groups ...*Group) *Manger {
	atomic.AddUint32((*uint32)(&m.level), 1)
	return m.addGroups(groups...)
}

// AddGroups Add Groups as normal
func (m *Manger) AddGroups(groups ...*Group) *Manger {
	return m.addGroups(groups...)
}

func (m *Manger) addGroups(groups ...*Group) *Manger {
	for _, group := range groups {
		group.SetLevel(m.level)
		m.groups = append(m.groups, group)
	}
	return m
}

func (m *Manger) Groups() []*Group {
	return m.groups
}
