package runner

import (
	"context"
	"time"

	"github.com/wuqinqiang/easycar/core/coordinator"
	"github.com/wuqinqiang/easycar/tools"

	"github.com/wuqinqiang/easycar/core/entity"

	"github.com/wuqinqiang/easycar/logging"

	"github.com/wuqinqiang/easycar/core/dao"

	"github.com/wuqinqiang/easycar/tools/retry"
)

type Runner struct {
	ticker  *time.Ticker
	backoff *retry.Retry
	cancel  func()
	ctx     context.Context

	dao         dao.TransactionDao
	coordinator *coordinator.Coordinator

	options *Options
}

func New(coordinator *coordinator.Coordinator, dao dao.TransactionDao, opts ...Option) *Runner {
	options := Default
	for _, fn := range opts {
		fn(options)
	}
	r := &Runner{
		coordinator: coordinator,
		dao:         dao,
		options:     options,
		backoff:     retry.New(10, retry.WithMaxBackOffTime(2*time.Minute), retry.WithFactor(2)),
	}
	r.ticker = time.NewTicker(r.options.duration)

	r.ctx, r.cancel = context.WithCancel(context.Background())
	go r.loop()
	return r
}

func (r *Runner) loop() {
	defer r.ticker.Stop()
	for {
		select {
		case <-r.ticker.C:
			list, err := r.dao.FindProcessingList(r.ctx, 2)
			if err != nil {
				logging.Errorf("[runner] loop err:%v", err)
				continue
			}

			if len(list) == 0 {
				// backoff for ticker
				duration := r.backoff.Duration()
				if duration != r.backoff.MaxBackOffTime() {
					// back off
					logging.Infof("[duration] Reset ticker:%v", duration)
					r.ticker.Reset(duration)
				}
				continue
			}
			r.runJob(list)
			// reset to default value
			r.ticker.Reset(r.options.duration)
			r.backoff.Reset()
		case <-r.ctx.Done():
			return
		}
	}
}

func (r *Runner) runJob(list []*entity.Global) {
	// default ctx
	ctx := context.Background()

	for i := 0; i < len(list); i++ {
		global := list[i]
		branches, err := r.coordinator.GetBranchList(ctx, global.GetGId())
		if err != nil {
			continue
		}
		tools.GoSafe(func() {
			var (
				err error
			)
			defer func() {
				if err != nil {
					logging.Errorf("[Runner] err:%v", err)
				}
			}()

			if global.Phase1() {
				err = r.coordinator.Phase1(ctx, global, branches)
			} else if global.Phase2() {
				err = r.coordinator.Phase2(ctx, global, branches)
			} else {
				logging.Warnf("[Runner] global:%v state :%v is wrong", global.GID, global.State)
			}
		})

	}
}

func (r *Runner) Run(ctx context.Context) (_ error) { return nil }

func (r *Runner) Stop(ctx context.Context) error {
	r.cancel()
	return nil
}
