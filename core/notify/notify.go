package notify

import (
	"context"

	"github.com/wuqinqiang/easycar/logging"
	"github.com/wuqinqiang/easycar/tools"
)

type Notify interface {
	Notify(content Content)
	Stop()
}

type Sender interface {
	Send(title, msg string) error
}

type notify struct {
	ctx     context.Context
	cancel  func()
	senders []Sender
	ch      chan Content
}

func New(senders []Sender) Notify {
	n := &notify{
		senders: senders,
		ch:      make(chan Content, 50),
	}
	n.ctx, n.cancel = context.WithCancel(context.Background())
	go n.waitEvent()
	return n
}

func (n *notify) Notify(content Content) {
	if len(n.senders) == 0 {
		return
	}
	n.ch <- content
}
func (n *notify) Stop() {
	n.cancel()
}

func (n *notify) waitEvent() {
	for {
		select {
		case <-n.ctx.Done():
			return
		case content := <-n.ch:
			for _, sender := range n.senders {
				tools.GoSafe(func() {
					err := sender.Send("easycar", content.Msg())
					if err != nil {
						logging.Errorf("[waitEvent]:%v", err)
					}
				})
			}
		}
	}
}
