package grpc

type Parser interface {
	Get() (service string, method string, err error)
}
