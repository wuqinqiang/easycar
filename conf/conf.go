package conf

import (
	"fmt"

	"github.com/wuqinqiang/easycar/conf/file"
)

var _ Config = (*config)(nil)

type Config interface {
	Load() error
	GetValueByKey(key string) (interface{}, bool)
}

type config struct {
	opts   options
	reader Reader
}

func New(optFunS ...OptionFunc) Config {
	// default options
	opt := options{
		decoder: DefaultDecoder,
		source:  file.NewFile(""),
	}

	for _, optFunc := range optFunS {
		optFunc(&opt)
	}
	return &config{opts: opt, reader: NewReader(opt)}
}

func (c *config) Load() error {
	if c.opts.source == nil {
		return fmt.Errorf("must be set source")
	}
	value, err := c.opts.source.Load()
	if err != nil {
		return err
	}
	err = c.reader.LoadValue(value)
	if err != nil {
		return err
	}
	return nil
}

func (c *config) GetValueByKey(key string) (interface{}, bool) {
	return c.reader.GetValue(key)
}
