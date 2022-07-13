package consts

type ( // GlobalState is a global state of the application
	GlobalState string
	// BranchState branch state, such as tcc (try、confirm、cancel)
	BranchState string
)

const (
	// Begin global state
	Begin GlobalState = "begin"

	//************global state***************

	GlobalCommitting     GlobalState = "committing"
	GlobalCommitted      GlobalState = "committed"
	GlobalCommitRetrying GlobalState = "commitRetrying"
	GlobalCommitFailed   GlobalState = "commitFailed"

	GlobalRollBacking      GlobalState = "rollBacking"
	GlobalRollBacked       GlobalState = "rollBacked"
	GlobalRollBackRetrying GlobalState = "rollBackRetrying"
	GlobalRollBackFailed   GlobalState = "rollBackFailed"

	//************branch state***************

	BranchReady     BranchState = "branchReady"
	BranchRetrying  BranchState = "branchRetrying"
	BranchSucceed   BranchState = "succeed"
	BranchFailState BranchState = "failed"
)
