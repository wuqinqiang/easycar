package client

type RM struct {
	// server Address
	serverAddress string
	// todo some timeout config
}

type TransactionInterface interface {
	GetTransactionName() string
}

func NewRM(serverAddress string) *RM {
	return &RM{serverAddress: serverAddress}
}

// Start start for a transaction
func (r *RM) Start() {

}

// RegisterBranch Register transaction branch to server
func (r *RM) RegisterBranch() {

}
