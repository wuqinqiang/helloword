package notify

import (
	"context"
	"fmt"
	"github.com/wuqinqiang/helloword/notify/base"

	"github.com/wuqinqiang/helloword/tools"
)

type Notify interface {
	Notify(subject base.Subject)
	Stop()
}

type Sender interface {
	Send(subject base.Subject) error
}

type notify struct {
	ctx     context.Context
	cancel  func()
	senders []Sender
	ch      chan base.Subject
}

func New(senders []Sender) Notify {
	n := &notify{
		senders: senders,
		ch:      make(chan base.Subject, 50),
	}
	n.ctx, n.cancel = context.WithCancel(context.Background())
	go n.waitEvent()
	return n
}

func (n *notify) Notify(subject base.Subject) {
	if len(n.senders) == 0 {
		return
	}
	n.ch <- subject
}
func (n *notify) Stop() {
	n.cancel()
}

func (n *notify) waitEvent() {
	for {
		select {
		case <-n.ctx.Done():
			return
		case subject := <-n.ch:
			for _, sender := range n.senders {
				tools.GoSafe(func() {
					err := sender.Send(subject)
					if err != nil {
						fmt.Println()
					}
				})
			}
		}
	}
}
