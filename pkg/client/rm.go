package client

type RM struct {
	// server Address
	serverAddress string
	// todo some timeout config
}

func NewRM() *RM {
	return &RM{}
}

// Start start for a transaction
func (r *RM) Start() {
}

// RegisterBranch Register transaction branch to server
func (r *RM) RegisterBranch() {

}
