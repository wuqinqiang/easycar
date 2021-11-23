package client

var _ TransactionInterface = &TCC{}

type TCCOption func(tcc *TCC)

var (
	DefaultProtoCol = "http"
)

type TCC struct {
	protoCol string
}

func NewTCC(options ...TCCOption) *TCC {
	tcc := &TCC{protoCol: DefaultProtoCol}
	for _, option := range options {
		option(tcc)
	}
	return tcc
}

func (t *TCC) GetTransactionName() string {
	return "tcc"
}

func (t *TCC) GetProtocol() string {
	return t.protoCol
}

func (t *TCC) RegisterBranch() {

}
