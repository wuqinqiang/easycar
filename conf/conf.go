package conf

import (
	"github.com/wuqinqiang/easycar/conf/common"
	"github.com/wuqinqiang/easycar/conf/file"
)

type (
	Mode string
)

const (
	File Mode = "file"
	Etcd Mode = "etcd"
	Env  Mode = "env"
	//Add more conf schema here
)

type Conf interface {
	Load() (*common.EasyCar, error)
}

func NewConf(mode string) (Conf, error) {
	m := Mode(mode)
	switch m {
	case Env:
	case Etcd:
	default:
	}
	return file.NewFile("/conf.yml"), nil
}
