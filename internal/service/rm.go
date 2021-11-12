package service

type RM struct {
}

// Begin  begin a new transaction, return globleId
func (rm RM) Begin() (gId string, err error) {
	panic("dc")
}

// Submit summit transaction
func (rm RM) Submit(gId string) {
	panic("dd")
}
