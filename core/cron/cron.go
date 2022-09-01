package cron

import (
	"context"

	"github.com/robfig/cron/v3"
)

type Job func(ctx context.Context)

var (
	runner *Runner
)

func init() {
	runner = &Runner{
		cron: cron.New(cron.WithParser(cron.NewParser(
			cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor))),
	}
}

func AddCronJob(spec string, job Job) (err error) {
	cronJob := cron.NewChain().Then(cron.FuncJob(func() {
		job(context.Background())
	}))
	_, err = runner.cron.AddJob(spec, cronJob)
	return
}

type Runner struct {
	ctx  context.Context
	cron *cron.Cron
}

func (r *Runner) Start(ctx context.Context) error {
	r.cron.Start()
	return nil
}

func (r *Runner) Stop(ctx context.Context) error {
	r.cron.Stop()
	return nil
}
