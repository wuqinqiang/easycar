package runner

import "time"

var Default = &Options{
	duration:     time.Second * 2,
	timeInterval: time.Minute * 5,
	MaxTimes:     3,
}

type Option func(options *Options)

type Options struct {
	duration     time.Duration
	MaxTimes     int
	timeInterval time.Duration
}

func WitDuration(duration time.Duration) Option {
	return func(options *Options) {
		options.duration = duration
	}
}

func WithMaxTimes(maxTimes int) Option {
	return func(options *Options) {
		if maxTimes > 0 {
			options.MaxTimes = maxTimes
		}
	}
}

func WithTimeInterval(timeInterval int) Option {
	return func(options *Options) {
		if timeInterval > 0 {
			options.timeInterval = time.Minute * time.Duration(timeInterval)
		}
	}
}
