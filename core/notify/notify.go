package notify

import (
	"context"

	"github.com/wuqinqiang/easycar/logging"
	"github.com/wuqinqiang/easycar/tools"
)

type Notify interface {
	Notify(content Content)
}

type Sender interface {
	Send(title, msg string) error
}

type notify struct {
	ctx     context.Context
	senders []Sender
	ch      chan Content
}

func New(ctx context.Context, senders []Sender) Notify {
	n := &notify{
		ctx:     ctx,
		senders: senders,
		ch:      make(chan Content, 50),
	}
	go n.waitEvent()
	return n
}

func (n *notify) Notify(content Content) {
	if len(n.senders) == 0 {
		return
	}
	n.ch <- content
}
func (n *notify) waitEvent() {
	logging.Infof("notify start")
	for {
		select {
		case <-n.ctx.Done():
			logging.Infof("Received the done signal")
			return
		case content := <-n.ch:
			_ = content
			for _, sender := range n.senders {
				tools.GoSafe(func() {
					err := sender.Send("easycar", "err")
					if err != nil {
						logging.Errorf("[waitEvent]:%v", err)
					}
				})
			}
		}
	}
}
