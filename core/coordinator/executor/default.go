package executor

import (
	"time"

	"github.com/wuqinqiang/easycar/core/transport"
)

// DefaultExecutor executor
var DefaultExecutor = &Default{
	manager: transport.NewManager(),
	// default timeout for branches
	timeout: 8 * time.Second,
}
