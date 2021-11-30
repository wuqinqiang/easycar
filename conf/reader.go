package conf

import "sync"

var _ Reader = (*reader)(nil)

type Reader interface {
	LoadValue(value KeyValue) error
	GetValue(key string) (val interface{}, ok bool)
}

type reader struct {
	opts   *options
	values map[string]interface{}
	lock   sync.Mutex
}

func NewReader(opts *options) Reader {
	return &reader{
		opts:   opts,
		values: make(map[string]interface{}),
		lock:   sync.Mutex{},
	}
}
func (r *reader) LoadValue(value KeyValue) error {
	// todo un data to values by keyValue
	return nil
}

func (r *reader) GetValue(key string) (val interface{}, ok bool) {
	// get value by key from values
	return nil, false
}
