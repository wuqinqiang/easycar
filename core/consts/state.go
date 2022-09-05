package consts

type ( // GlobalState is a global state of the application
	GlobalState string
	// BranchState branch state, such as tcc (try、confirm、cancel)
	BranchState string
)

const (
	// Init init global state
	Init GlobalState = "init"

	//************global state***************
	Phase1Processing GlobalState = "phase1_processing"
	Phase1Retrying   GlobalState = "phase1_retrying"
	Phase1Failed     GlobalState = "phase1_failed"
	Phase1Success    GlobalState = "phase1_success"

	Phase2Committing     GlobalState = "phase2_committing"
	Phase2Rollbacking    GlobalState = "phase2_rollbacking"
	Phase2CommitFailed   GlobalState = "phase2_commit_failed"
	Phase2RollbackFailed GlobalState = "phase2_rollback_failed"

	//Committed Distributed transaction executed successfully
	Committed  GlobalState = "committed"
	Rollbacked GlobalState = "rollbacked"

	//************branch state***************

	BranchReady     BranchState = "ready"
	BranchRetrying  BranchState = "retrying"
	BranchSucceed   BranchState = "succeed"
	BranchFailState BranchState = "failed"
)
