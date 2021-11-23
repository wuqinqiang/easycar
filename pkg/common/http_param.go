package common

import "strconv"

type GlobalData struct {
	transactionName TransactionName
	protocol        string
}

func (g *GlobalData) SetTransactionName(name TransactionName) {
	g.transactionName = name
}

func (g *GlobalData) GetTransactionName() TransactionName {
	return g.transactionName
}

func (g *GlobalData) SetProtocol(protocol string) {
	g.protocol = protocol
}

func (g *GlobalData) GetProtocol() string {
	return g.protocol
}

// BranchData client transaction data
type BranchData struct {
	// global id
	gId string
	// rm request url
	url string
	// request rm for data
	reqData interface{}
	// branchId
	branId     string
	branchType string
}

func (b *BranchData) SetGId(gId string) {
	b.gId = gId
}

func (b *BranchData) GetGid() string {
	return b.gId
}

func (b *BranchData) SetUrl(url string) {
	b.url = url
}

func (b *BranchData) GetUrl() string {
	return b.url
}

func (b *BranchData) SetReqData(data interface{}) {
	b.reqData = data
}

func (b *BranchData) SetBranId() {
	b.branId = b.gId + strconv.Itoa(1)
}
