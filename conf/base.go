package conf

type Config interface {
	Load() error
	GetValueByKey(key string) interface{}
}

type config struct {
	reader Reader
}

func New(optFuns ...OptionFunc) Config {
	// default options
	opt := &options{}
	for _, optFunc := range optFuns {
		optFunc(opt)
	}
	return &config{reader: NewReader(opt)}
}

func (c config) Load() error {
	panic("implement me")
}

func (c config) GetValueByKey(key string) interface{} {
	panic("implement me")
}
