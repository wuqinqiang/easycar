package runner

import (
	"context"
	"sync"

	"github.com/wuqinqiang/easycar/logging"

	"github.com/robfig/cron/v3"
)

type Job func(ctx context.Context)

func NewRunner(spec string, job Job) (*Runner, error) {
	runner := &Runner{
		once: sync.Once{},
		cron: cron.New(cron.WithParser(cron.NewParser(
			cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor))),
	}
	_, err := runner.cron.AddJob(spec, cron.NewChain().Then(cron.FuncJob(func() {
		job(context.Background())
	})))
	if err != nil {
		return nil, err
	}
	return runner, nil
}

type Runner struct {
	cron *cron.Cron
	once sync.Once
}

func (r *Runner) Run(ctx context.Context) error {
	r.cron.Start()
	logging.Info("[Runner] start")
	return nil
}

func (r *Runner) Stop(ctx context.Context) error {
	r.once.Do(func() {
		r.cron.Stop()
	})
	logging.Info("[Runner] stopped")
	return nil
}
