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
	Phase1Retrying  GlobalState = "p1_retrying"
	Phase1Failed    GlobalState = "p1_failed"
	Phase1Success   GlobalState = "p1_success"

	//**********Phase2*************
	Phase2Committing     GlobalState = "p2_committing"
	Phase2CommitFailed   GlobalState = "p2_commit_failed"
	Phase2CommitRetrying GlobalState = "p2_commit_retrying"

	Phase2Rollbacking      GlobalState = "p2_rollbacking"
	Phase2RollbackFailed   GlobalState = "p2_rollback_failed"
	Phase2RollbackRetrying GlobalState = "p2_rollback_retrying"

	Committed  GlobalState = "committed"
	Rollbacked GlobalState = "rollbacked"

	//************branch state***************

	BranchInit      BranchState = "init"
	BranchRetrying  BranchState = "retrying"
	BranchSucceed   BranchState = "succeed"
	BranchFailState BranchState = "failed"
)
