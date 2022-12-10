package client

import (
	"strings"

	"github.com/wuqinqiang/easycar/core/consts"

	"github.com/wuqinqiang/easycar/proto"
)

type (
	Branch struct {
		// uri branch request uri.
		// http://127.0.0.1:50062/account/confirmDebit
		// grpc:127.0.0.1:50061/stock.Stock/TryDeduct
		uri string
		//data the branch requested data
		data []byte
		//data the branch requested header,if protocol is grpc,it is empty
		header []byte
		// branch action.
		//etc. TCC(try、confirm、cancel)
		action consts.BranchAction
		// level as a branch stratification basis,the smaller the number,the earlier phase 1 is executed
		level consts.Level
		// timeout for request branch.unit of second
		timeout int64
		// protocol branch network protocol. only have HTTP/gRPC
		protocol Protocol
	}
)

func NewBranch(uri string, action consts.BranchAction) *Branch {
	b := &Branch{
		action:  action,
		level:   1,
		timeout: 3, // default
	}
	b.uri, b.protocol = protocol(uri)
	return b
}

// SetLevel set branch currentLevel
func (branch *Branch) SetLevel(level consts.Level) *Branch {
	if level == 0 {
		return branch
	}
	branch.level = level
	return branch
}

// SetProtocol set branch network protocol
func (branch *Branch) SetProtocol(protocol Protocol) *Branch {
	branch.protocol = protocol
	return branch
}

// SetData set branch request data
func (branch *Branch) SetData(data []byte) *Branch {
	branch.data = data
	return branch
}

// SetHeader set branch header
func (branch *Branch) SetHeader(header []byte) *Branch {
	branch.header = header
	return branch
}

// SetTimeout set timeout for request branch
func (branch *Branch) SetTimeout(timeout int64) *Branch {
	branch.timeout = timeout
	return branch
}

// Convert from client's Branch to pb.RegisterReq_Branch
func (branch *Branch) Convert() *proto.RegisterReq_Branch {
	var (
		req proto.RegisterReq_Branch
	)
	req.Uri = branch.uri
	req.ReqData = string(branch.data)
	req.ReqHeader = string(branch.header)
	req.Protocol = string(branch.protocol)
	req.Timeout = int32(branch.timeout)
	req.Action = consts.ConvertBranchActionToGrpc(branch.action)
	req.Level = int32(branch.level)
	return &req
}

func protocol(uri string) (string, Protocol) {
	if strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://") {
		return uri, HTTP
	}

	if strings.HasPrefix(uri, "grpc://") {
		return uri, GRPC
	}
	// 127.0.0.1:50060/order.Order/Cancel
	keys := strings.Split(uri, "//")
	if len(keys) == 1 && keys[0] == uri {
		return "grpc://" + uri, GRPC
	}
	return uri, Undefined
}
