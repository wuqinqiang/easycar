package executor

import (
	"github.com/wuqinqiang/easycar/core/transport"
)

var DefaultExecutor = &executor{
	manager: transport.NewManager(),
}
