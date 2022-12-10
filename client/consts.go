package client

const (
	HTTP      Protocol = "http"
	GRPC      Protocol = "grpc"
	Undefined Protocol = "undefined"
)

type (
	Protocol string
)
