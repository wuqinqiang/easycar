package common

type (
	GlobalState     string
	TransactionName string
)

const (
	TCC  TransactionName = "tcc"
	SAGA TransactionName = "saga"

	PreparedState GlobalState = "prepared"
	SucceedState  GlobalState = "succeed"
	FailedState   GlobalState = "failed"
)
