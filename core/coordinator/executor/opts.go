package executor

import "github.com/wuqinqiang/easycar/core/transport"

var DefaultExecutor = &executor{
	option: &Option{
		openRetry: false,
	},
	manager: transport.NewManager(),
}

type (
	OptFn func(opt *Option)

	Option struct {
		allowRetries uint32
		factor       uint32
		openRetry    bool
	}
)
