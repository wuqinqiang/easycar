package entity

type (
	BranchType  uint8
	BranchState string
)

const (
	BranchReadyState    = "ready"
	BranchSucceedState  = "succeed"
	BranchFinishedState = "fail"
)

type Branch struct {
	gId        string
	url        string
	reqData    string
	branchId   string
	branchType BranchType
	state      BranchState
}

func NewBranch(gId string) *Branch {
	return &Branch{
		gId: gId,
	}
}

func (b *Branch) SetGId(gId string) {
	b.gId = gId
}
func (b *Branch) GetGId() string {
	return b.gId
}

func (b *Branch) SetUrl(url string) {
	b.url = url
}
func (b *Branch) GetUrl() string {
	return b.url
}

func (b *Branch) SetReqData(reqData string) {
	b.reqData = reqData
}
func (b *Branch) GetReqData() string {
	return b.reqData
}

func (b *Branch) SetBranchId(branchId string) {
	b.branchId = branchId
}
func (b *Branch) GetBranchId() string {
	return b.branchId
}

func (b *Branch) SetBranchType(branchType BranchType) {
	b.branchType = branchType
}
func (b *Branch) GetBranchType() BranchType {
	return b.branchType
}

func (b *Branch) Setstate(state BranchState) {
	b.state = state
}
func (b *Branch) GetBranchState() BranchState {
	return b.state
}

func (b *Branch) CanHandle() bool {
	return !(b.GetBranchState() == BranchSucceedState || b.GetBranchState() == BranchFinishedState)
}
