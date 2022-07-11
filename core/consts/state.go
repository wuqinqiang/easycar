package consts

type ( // GlobalState is a global state of the application
	GlobalState string
	// BranchState branch state, such as tcc (try、confirm、cancel)
	BranchState string
)

const (
	// Begin global state
	Begin GlobalState = "begin"

	Submitting    GlobalState = "submitting"
	Submitted     GlobalState = "submitted"
	SubmitFailed  GlobalState = "commitFailed"
	TimeOutSubmit GlobalState = "timeOutSubmit"

	RollBacking       GlobalState = "rollBacking"
	RollBacked        GlobalState = "rollBacked"
	RollBackFailed    GlobalState = "rollBackFailed"
	TimeoutRollBacked GlobalState = "timeoutRollBacked"

	UNKNOWN GlobalState = "unknown"
	// end global state

	BranchReadyState   BranchState = "ready"
	BranchSucceedState BranchState = "succeed"
	BranchFailState    BranchState = "fail"
)
