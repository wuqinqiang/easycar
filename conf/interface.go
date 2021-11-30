package conf

type KeyValue struct {
	Value []byte
}

type Source interface {
	Load() (*KeyValue, error)
}
