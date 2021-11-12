package service

type GlobalState uint8

type Global struct {
	gId        string
	branchType uint8
	state      GlobalState
	protocol   string
}

//func NewGlobal(GId string, branchType) {
//
//}
