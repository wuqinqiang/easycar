package entity

import (
	"github.com/wuqinqiang/easycar/core/consts"
)

type (
	Branch struct {
		GID      string                 `gorm:"column:g_id;type:varchar(255);not null"`                      // global id
		BranchId string                 `gorm:"column:branch_id;type:varchar(255);not null"`                 // branch id
		Url      string                 `gorm:"column:url;type:varchar(255);not null"`                       // branch request url (example grpc or http)
		ReqData  string                 `gorm:"column:req_data;type:varchar(255);not null"`                  // request data
		TranType consts.TransactionType `gorm:"column:tran_type;type:varchar(255);not null"`                 // transaction type:tcc or saga or others
		PId      string                 `gorm:"column:p_id;type:varchar(255);not null"`                      // parent branch id
		Protocol string                 `gorm:"column:protocol;type:varchar(255);not null;default:http"`     //http,grpc
		Action   consts.BranchAction    `gorm:"column:action;type:varchar(255);not null"`                    // action type of transaction
		State    consts.BranchState     `gorm:"column:state;type:varchar(255);not null;default:branchReady"` // branch State
		//ChildrenList      []*Branch               //	children branch list
		EndTime int64 `gorm:"column:end_time;type:int;not null;default:0"`
		// 07-10 add
		Level consts.Level `gorm:"column:level;type:int;not null;default:1"` // branch level in tree

		LastErrMsg string `gorm:"column:last_err_msg;type:varchar(255);not null"`
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
