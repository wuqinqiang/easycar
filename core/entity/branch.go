package entity

import (
	"github.com/wuqinqiang/easycar/core/consts"
)

type (
	Branch struct {
		GID        string                 `gorm:"column:g_id;not null" bson:"g_id"`                      // global id
		BranchId   string                 `gorm:"column:branch_id;not null" bson:"branch_id"`            // branch id
		Url        string                 `gorm:"column:url;not null" bson:"url"`                        // branch request url (example grpc or http)
		ReqData    string                 `gorm:"column:req_data;not null" bson:"req_data"`              // request data
		ReqHeader  string                 `gorm:"column:req_header;not null" bson:"req_header"`          // request data
		TranType   consts.TransactionType `gorm:"column:tran_type;not null" bson:"tran_type"`            // transaction type:tcc or saga or others
		Protocol   string                 `gorm:"column:protocol;not null;default:http" bson:"protocol"` //http,grpc
		Action     consts.BranchAction    `gorm:"column:action;not null" bson:"action"`                  // action type of transaction
		State      consts.BranchState     `gorm:"column:state;not null;default:init" bson:"state"`       // branch State
		Level      consts.Level           `gorm:"column:level;not null;default:1" bson:"level"`          // branch level in tree
		LastErrMsg string                 `gorm:"column:last_err_msg;not null" bson:"last_err_msg"`
		Timeout    int64                  `gorm:"column:timeout;not null;default:0" bson:"timeout"`                 //request branch timeout(seconds)
		CreateTime int64                  `gorm:"create_time;autoCreateTime" json:"create_time" bson:"create_time"` // create time
		UpdateTime int64                  `gorm:"update_time;autoCreateTime" json:"update_time" bson:"update_time"` // last update time
		// todo add group id for branches
	}
	BranchList []*Branch
)

func (b Branch) TableName() string {
	return "branch"
}

func (b *Branch) IsTcc() bool {
	return b.TranType == consts.TCC
}
func (b *Branch) TccTry() bool {
	return b.Action == consts.Try && b.IsTcc()
}

func (b *Branch) TccCancel() bool {
	return b.Action == consts.Cancel && b.IsTcc()
}

func (b *Branch) TccConfirm() bool {
	return b.Action == consts.Confirm
}

func (b *Branch) SAGA() bool {
	return b.TranType == consts.SAGA
}

func (b *Branch) Success() bool {
	return b.State == consts.BranchSucceed
}

func (b *Branch) SAGANormal() bool {
	return b.Action == consts.Normal && b.SAGA()
}

func (b *Branch) SAGACompensation() bool {
	return b.Action == consts.Compensation && b.SAGA()
}
