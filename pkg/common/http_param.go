package common

import "strconv"

// ReqGlobalData begin a transaction
type ReqGlobalData struct {
	transactionName TransactionName
	protocol        string
}

func (g *ReqGlobalData) SetTransactionName(name TransactionName) {
	g.transactionName = name
}

func (g *ReqGlobalData) GetTransactionName() TransactionName {
	return g.transactionName
}

func (g *ReqGlobalData) SetProtocol(protocol string) {
	g.protocol = protocol
}

func (g *ReqGlobalData) GetProtocol() string {
	return g.protocol
}

// RespBase base resp form transaction manager server
type RespBase struct {
	Msg string
	Err error
}

func (respBase RespBase) GetMsg() string {
	return respBase.Msg
}
func (respBase RespBase) GetError() error {
	return respBase.Err
}

// RespGlobalData begin transaction resp
type RespGlobalData struct {
	RespBase
	GId string
}

func (globalResp RespGlobalData) GetGId() string {
	return globalResp.GId
}

//BranchData registerBranch req data
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

func (b *BranchData) GetReqData() interface{} {
	return b.GetReqData()
}

func (b *BranchData) GetBranId() {
	b.branId = b.gId + strconv.Itoa(1)
}

// ReportStateData report tm that transaction state
type ReportStateData struct {
	GId string
}

func (b *ReportStateData) SetGId(gId string) {
	b.GId = gId
}
