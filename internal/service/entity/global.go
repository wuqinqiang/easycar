package entity

type (
	GlobalState     uint8
	TransactionType string
)

const (
	TCC  TransactionType = "tcc"
	SAGA TransactionType = "saga"

	PreparedState GlobalState = iota + 1
	SubmittedState
	AbortingState
	RollbackState
)

type Global struct {
	gId             string
	transactionName TransactionType
	state           GlobalState
	protocol        string
}

func NewGlobal(gId string) *Global {
	return &Global{
		gId: gId,
	}
}

func (g *Global) SetGId(gId string) {
	g.gId = gId
}

func (g *Global) GetGId() string {
	return g.gId
}

func (g *Global) SetBranchType(transactionName TransactionType) {
	g.transactionName = transactionName
}

func (g *Global) GetBranchType() TransactionType {
	return g.transactionName
}

func (g *Global) SetState(state GlobalState) {
	g.state = state
}

func (g *Global) GetState() GlobalState {
	return g.state
}

func (g *Global) SetProtocol(protocol string) {
	g.protocol = protocol
}

func (g *Global) GetProtocol() string {
	return g.protocol
}
