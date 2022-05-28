package mode

import "github.com/wuqinqiang/easycar/core/entity"

type Tcc struct {
}

func NewTcc() Tcc {
	return Tcc{}
}

func (t Tcc) Mode() string {
	return "tcc"
}

func (t Tcc) prepare(branch entity.Branch) error {
	return nil
}

func (t Tcc) commit(branch entity.Branch) error {
	return nil
}

func (t Tcc) rollback(branch entity.Branch) error {
	return nil
}
