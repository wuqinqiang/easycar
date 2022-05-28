package mode

import "github.com/wuqinqiang/easycar/core/entity"

type Saga struct {
}

func NewSaga() Saga {
	return Saga{}
}

func (saga Saga) Mode() string {
	return "saga"
}

func (saga Saga) normalOperation(branch entity.Branch) error {
	return nil
}

func (saga Saga) compensating(branch entity.Branch) error {
	return nil
}
