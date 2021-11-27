package common

const ()

type (
	GlobalState     string
	TransactionName string

	// BranchType branch action,such as tcc (try、confirm、cancel)
	BranchType  string
	BranchState string
)

const (
	TCC  TransactionName = "tcc"
	SAGA TransactionName = "saga"

	Prepared  GlobalState = "prepared"
	Submitted GlobalState = "submitted"
	Succeed   GlobalState = "succeed"
	Failed    GlobalState = "failed"
	Abort     GlobalState = "abort"

	Try     BranchType = "try"
	Confirm BranchType = "confirm"
	Cancel  BranchType = "cancel"

	BranchReadyState    BranchState = "ready"
	BranchSucceedState  BranchState = "succeed"
	BranchFinishedState BranchState = "fail"
)
