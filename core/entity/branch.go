package entity

import (
	"github.com/wuqinqiang/easycar/core/consts"
)

type (
	Branch struct {
		GID        string                 `gorm:"column:g_id;not null"`                  // global id
		BranchId   string                 `gorm:"column:branch_id;not null"`             // branch id
		Url        string                 `gorm:"column:url;not null"`                   // branch request url (example grpc or http)
		ReqData    string                 `gorm:"column:req_data;not null"`              // request data
		ReqHeader  string                 `gorm:"column:req_header;not null"`            // request data
		TranType   consts.TransactionType `gorm:"column:tran_type;not null"`             // transaction type:tcc or saga or others
		Protocol   string                 `gorm:"column:protocol;not null;default:http"` //http,grpc
		Action     consts.BranchAction    `gorm:"column:action;not null"`                // action type of transaction
		State      consts.BranchState     `gorm:"column:state;not null;default:ready"`   // branch State
		Level      consts.Level           `gorm:"column:level;not null;default:1"`       // branch level in tree
		LastErrMsg string                 `gorm:"column:last_err_msg;not null"`
		Timeout    int64                  `gorm:"column:timeout;not null;default:0"`             //request branch timeout(seconds)
		CreateTime int64                  `gorm:"create_time;autoCreateTime" json:"create_time"` // create time
		UpdateTime int64                  `gorm:"update_time;autoCreateTime" json:"update_time"` // last update time
	}
	BranchList []*Branch
)

func (b Branch) TableName() string {
	return "branch"
}

func (b *Branch) IsSucceed() bool {
	return b.State == consts.BranchSucceed
}

func (b *Branch) IsBranchFailState() {

}

func (b *Branch) IsTcc() bool {
	return b.TranType == consts.TCC
}
func (b *Branch) IsTccTry() bool {
	return b.Action == consts.Try && b.IsTcc()
}

func (b *Branch) IsTccCancel() bool {
	return b.Action == consts.Cancel && b.IsTcc()
}

func (b *Branch) IsTccConfirm() bool {
	return b.Action == consts.Confirm
}

func (b *Branch) CanDo() {

}

func (b *Branch) IsSAGA() bool {
	return b.TranType == consts.SAGA
}

func (b *Branch) IsSAGANormal() bool {
	return b.Action == consts.Normal && b.IsSAGA()
}

func (b *Branch) IsSAGACompensation() bool {
	return b.Action == consts.Compensation && b.IsSAGA()
}
