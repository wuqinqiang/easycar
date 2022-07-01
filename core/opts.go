package core

var (
	defaultOpts = opts{
		poolSize: 300,
	}
)

type opts struct {
	poolSize uint32
}

type OptFn func(*opts)

func WithPoolSize(size uint32) OptFn {
	return func(o *opts) {
		if size > 0 {
			o.poolSize = size
		}
	}
}
