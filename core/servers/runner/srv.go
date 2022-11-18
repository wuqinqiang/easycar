package runner

import (
	"context"
	"fmt"
	"sync"

	"github.com/wuqinqiang/easycar/tools"

	"github.com/wuqinqiang/easycar/core/coordinator"

	"github.com/wuqinqiang/easycar/core/dao"

	"github.com/wuqinqiang/easycar/logging"

	"github.com/robfig/cron/v3"
)

type Job func(ctx context.Context)

func NewRunner(spec string, c *coordinator.Coordinator) *Runner {
	runner := &Runner{
		dao:         dao.GetTransaction(),
		coordinator: c,
		cron: cron.New(cron.WithParser(cron.NewParser(
			cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor))),
	}
	if _, err := runner.cron.AddJob(spec, cron.NewChain().Then(cron.FuncJob(func() {
		runner.job()
	}))); err != nil {
		panic(err)
	}
	return runner
}

type Runner struct {
	coordinator *coordinator.Coordinator
	cron        *cron.Cron
	once        sync.Once
	dao         dao.TransactionDao
}

func (r *Runner) Run(ctx context.Context) error {
	r.cron.Start()
	logging.Infof("[Runner] start")
	return nil
}

func (r *Runner) Stop(ctx context.Context) error {
	r.once.Do(func() {
		r.cron.Stop()
	})
	logging.Infof("[Runner] stopped")
	return nil
}

func (r *Runner) job() {
	list, err := r.dao.FindProcessingList(context.Background(), 10)
	if err != nil {
		logging.Info(fmt.Sprintf("[job] err:%v", err))
		return
	}
	if len(list) == 0 {
		return
	}
	logging.Info(fmt.Sprintf("[processing list ]:%v", list))

	for i := 0; i < len(list); i++ {
		global := list[i]
		ctx := context.Background()
		branches, err := r.coordinator.GetBranchList(ctx, global.GetGId())
		if err != nil {
			continue
		}

		tools.GoSafe(func() {
			ctx := context.Background()
			var (
				err error
			)
			defer func() {
				if err != nil {
					logging.Error(fmt.Sprintf("[Runner] err:%v", err))
				}
			}()
			if global.Phase1Retrying() {
				err = r.coordinator.Phase1(ctx, global, branches)
			} else if global.Phase2Committing() || global.Phase2RollBacking() {
				err = r.coordinator.Phase2(ctx, global, branches)
			}
		})
	}

	fmt.Println("hello world")
}
