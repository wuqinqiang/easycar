package conf

type Decoder func(value *KeyValue, res map[string]interface{}) error

type OptionFunc func(*options)

type options struct {
	decoder Decoder
}

func WithDecoder(decoder Decoder) OptionFunc {
	return func(o *options) {
		o.decoder = decoder
	}
}
