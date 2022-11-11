package entity

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/wuqinqiang/easycar/core/consts"
	"github.com/wuqinqiang/easycar/proto"
)

func GetTranTypeByPb(protoType proto.TranType) consts.TransactionType {
	switch protoType {
	case proto.TranType_TCC:
		return consts.TCC
	case proto.TranType_SAGE:
		return consts.SAGA
	default:
	}
	return consts.TransactionUnknown
}

func GetActionByPb(action proto.Action) consts.BranchAction {
	switch action {
	case proto.Action_TRY:
		return consts.Try
	case proto.Action_CONFIRM:
		return consts.Confirm
	case proto.Action_CANCEL:
		return consts.Cancel
	case proto.Action_NORMAL:
		return consts.Normal
	case proto.Action_COMPENSATION:
		return consts.Compensation
	default:
	}
	return consts.ActionUnknown
}

// AssignmentByPb todo
func (b *Branch) AssignmentByPb(m *proto.RegisterReq_Branch) *Branch {
	b.Url = m.GetUri()
	b.ReqData = m.GetReqData()
	b.ReqHeader = m.GetReqHeader()
	b.TranType = GetTranTypeByPb(m.GetTranType())
	b.Protocol = m.GetProtocol()
	b.Action = GetActionByPb(m.GetAction())
	b.Level = consts.Level(m.GetLevel())
	b.State = consts.BranchInit
	now := time.Now().Unix()
	b.CreateTime = now
	b.UpdateTime = now
	b.Timeout = int64(m.Timeout)

	var (
		buffer bytes.Buffer
	)
	// such as:saga-normal-166210685961363520
	buffer.WriteString(string(b.TranType) + "-" + string(b.Action) + "-" + strconv.Itoa(int(time.Now().Unix())) +
		CreateCaptcha())
	b.BranchId = buffer.String()
	return b
}

func CreateCaptcha() string {
	return fmt.Sprintf("%07d", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(99999999))
}

func GetBranchList(gid string, mList []*proto.RegisterReq_Branch) (list BranchList) {
	for i := range mList {
		var (
			b Branch
		)
		b.GID = gid
		list = append(list, b.AssignmentByPb(mList[i]))
	}
	return
}
