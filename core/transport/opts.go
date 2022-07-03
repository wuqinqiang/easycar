package transport

type (
	Opts struct {
		body    []byte
		headers map[string]string
	}

	OptsFn func(*Opts)
)

func WithNetBody(b []byte) OptsFn {
	return func(netOpts *Opts) {
		netOpts.body = b
	}
}

func WithNetHeaders(heads map[string]string) OptsFn {
	return func(netOpts *Opts) {
		netOpts.headers = heads
	}
}
