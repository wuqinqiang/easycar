package consts

type (
	// BranchAction branch action,such as tcc (try、confirm、cancel)
	BranchAction string
)

const (
	// Try Try、Confirm and Cancel branch type for TCC
	Try     BranchAction = "try"
	Confirm BranchAction = "confirm"
	Cancel  BranchAction = "cancel"

	// Normal and Compensation branch type for SAGA
	Normal       BranchAction = "normal"
	Compensation BranchAction = "compensation"
)
