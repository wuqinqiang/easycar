package client

import "github.com/wuqinqiang/easycar/core/consts"

type Group struct {
	//tranType for groups
	tranType consts.TransactionType
	branches []*Branch
}

// NewTccGroup create a set of branches for TCC mode
func NewTccGroup(tryUri, confirmUri, cancelUri string) *Group {
	g := &Group{
		tranType: consts.TCC,
	}
	// timeout?
	g.branches = []*Branch{
		NewBranch(tryUri, consts.Try),
		NewBranch(confirmUri, consts.Confirm),
		NewBranch(cancelUri, consts.Cancel),
	}
	//g.SetProtocol(g.protocol)
	return g
}

// NewSagaGroup create a set of branches for Saga mode
func NewSagaGroup(normalUri, compensation string) *Group {
	g := &Group{
		tranType: consts.SAGA,
	}
	g.branches = []*Branch{
		NewBranch(normalUri, consts.Normal),
		NewBranch(compensation, consts.Compensation),
	}
	return g
}

func (g *Group) GetTranType() consts.TransactionType {
	return g.tranType
}

func (g *Group) SetData(data []byte) *Group {
	g.set(func(branch *Branch) {
		branch.SetData(data)
	})
	return g
}

func (g *Group) SetTimeout(second int) *Group {
	g.set(func(branch *Branch) {
		branch.timeout = int64(second)
	})
	return g
}

func (g *Group) SetHeader(data []byte) *Group {
	g.set(func(branch *Branch) {
		branch.SetHeader(data)
	})
	return g
}

func (g *Group) SetLevel(level consts.Level) *Group {
	g.set(func(branch *Branch) {
		branch.SetLevel(level)
	})
	return g
}

func (g *Group) set(fn func(branch *Branch)) {
	for _, branch := range g.branches {
		fn(branch)
	}
}
