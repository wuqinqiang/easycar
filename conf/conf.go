package conf

import (
	"github.com/wuqinqiang/easycar/conf/common"
	"github.com/wuqinqiang/easycar/conf/file"
)

type Conf interface {
	Load() (*common.EasyCar, error)
}

type (
	Mode string
)

const (
	File Mode = "file"
	Etcd Mode = "etcd"
	Env  Mode = "env"
	//Add more conf schema here
)

func NewConf(mode string) (Conf, error) {
	m := Mode(mode)
	switch m {
	case Env:
	case Etcd:
	default:
	}
	return file.NewFile("/conf.yml"), nil
}
