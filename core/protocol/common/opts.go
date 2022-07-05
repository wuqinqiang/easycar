package common

type (
	Opts struct {
		Body    []byte
		Headers map[string]string
	}

	OptsFn func(*Opts)
)

func WithBody(b []byte) OptsFn {
	return func(netOpts *Opts) {
		netOpts.Body = b
	}
}

func WithHeaders(heads map[string]string) OptsFn {
	return func(netOpts *Opts) {
		netOpts.Headers = heads
	}
}

func AppendHeaders(headers map[string]string) OptsFn {
	return func(opts *Opts) {
		for k, v := range headers {
			opts.Headers[k] = v
		}
	}
}
