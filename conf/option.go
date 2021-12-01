package conf

import (
	"fmt"

	"github.com/wuqinqiang/easycar/conf/encoding"
)

type Decoder func(value *KeyValue, res map[string]interface{}) error

type OptionFunc func(*options)

var (
	NotFoundDecoder = fmt.Errorf("not found decoder")
)

type options struct {
	decoder Decoder
	source  Source
}

func WithDecoder(decoder Decoder) OptionFunc {
	return func(o *options) {
		o.decoder = decoder
	}
}

func WithSource(source Source) OptionFunc {
	return func(o *options) {
		o.source = source
	}
}

func DefaultDecoder(res *KeyValue, target map[string]interface{}) error {
	if res.Format == "" {
		return NotFoundDecoder
	}
	codec, ok := encoding.GetCodecByFormat(res.Format)
	if !ok {
		return NotFoundDecoder
	}
	err := codec.UnMarshal(res.Value, target)
	if err != nil {
		return err
	}
	return nil
}
