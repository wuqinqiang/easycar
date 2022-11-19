package runner

import "time"

var Default = &Options{duration: time.Second * 2}

type Option func(options *Options)

type Options struct {
	duration time.Duration
	// todo add more options
}

func WitDuration(duration time.Duration) Option {
	return func(options *Options) {
		options.duration = duration
	}
}
