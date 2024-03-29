package consts

type ( // GlobalState is a global state of the application
	GlobalState string
	// BranchState branch state, such as tcc (try、confirm、cancel)
	BranchState string
)

const (
	//************global state***************

	// Init init global state
	Init GlobalState = "init"

	//**********Phase1************
	Phase1Preparing GlobalState = "preparing"
	Phase1Failed    GlobalState = "p1_failed"
	Phase1Success   GlobalState = "p1_success"

	//**********Phase2*************
	Phase2Committing   GlobalState = "p2_committing"
	Phase2CommitFailed GlobalState = "p2_commit_failed"

	Phase2Rollbacking    GlobalState = "p2_rollbacking"
	Phase2RollbackFailed GlobalState = "p2_rollback_failed"

	Committed  GlobalState = "committed"
	Rollbacked GlobalState = "rollbacked"

	//************branch state***************

	BranchInit      BranchState = "init"
	BranchRetrying  BranchState = "retrying"
	BranchSucceed   BranchState = "succeed"
	BranchFailState BranchState = "failed"
)

var P1InProgressStates = []string{
	string(Phase1Preparing),
}

var P2InProgressStates = []string{
	string(Phase1Success),
	string(Phase2Committing),
	string(Phase2CommitFailed),

	string(Phase1Failed),
	string(Phase2Rollbacking),
	string(Phase2RollbackFailed),
}
