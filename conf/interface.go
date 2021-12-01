package conf

type KeyValue struct {
	Value  []byte
	Format string
}

type Source interface {
	Load() (*KeyValue, error)
}
