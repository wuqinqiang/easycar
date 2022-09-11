package consts

type (
	// TransactionType 	transaction type
	TransactionType string
	// BranchAction branch action,such as tcc (try、confirm、cancel)
	BranchAction string

	// Level branch level,	0 is first level,same level can executed concurrently
	Level uint8
)

const (
	TransactionUnknown TransactionType = "unknown"
	TCC                TransactionType = "tcc"
	SAGA               TransactionType = "saga"

	// Try Try、Confirm and Cancel branch type for TCC

	ActionUnknown              = "unknown"
	Try           BranchAction = "try"
	Confirm       BranchAction = "confirm"
	Cancel        BranchAction = "cancel"

	// Normal and Compensation branch type for SAGA
	Normal       BranchAction = "normal"
	Compensation BranchAction = "compensation"
)
