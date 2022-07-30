package consts

type ( // GlobalState is a global state of the application
	GlobalState string
	// BranchState branch state, such as tcc (try、confirm、cancel)
	BranchState string
)

const (
	// Ready global state
	Ready GlobalState = "ready"

	//************global state***************
	Phase1Processing GlobalState = "phase1_processing"
	Phase1Retrying   GlobalState = "phase1_retrying"
	Phase1Failed     GlobalState = "phase1_failed"
	Phase1Success    GlobalState = "phase1_success"

	Phase2Processing GlobalState = "phase2_processing"
	Phase2Retrying   GlobalState = "phase2_retrying"
	Phase2Failed     GlobalState = "phase2_failed"
	Phase2Success    GlobalState = "phase2_success"

	//************branch state***************

	BranchReady     BranchState = "branchReady"
	BranchRetrying  BranchState = "branchRetrying"
	BranchSucceed   BranchState = "succeed"
	BranchFailState BranchState = "failed"
)
