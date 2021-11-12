package service

type (
	BranchType  uint8
	BranchState uint8
)

type Branch struct {
	gId        string
	url        string
	reqData    string
	branchId   string
	branchType BranchType
	state      BranchState
}
