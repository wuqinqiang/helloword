package notify

import (
	"context"
	"fmt"
	"sync"

	"github.com/wuqinqiang/helloword/notify/base"

	"github.com/wuqinqiang/helloword/tools"
)

type Notify interface {
	Notify(subject base.Subject)
	Stop()
	Wait()
}

type Sender interface {
	Send(subject base.Subject) error
}

type notify struct {
	ctx     context.Context
	cancel  func()
	senders []Sender
	ch      chan base.Subject
	wg      sync.WaitGroup
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
	n.wg.Add(len(n.senders))
}

func (n *notify) Stop() {
	n.cancel()
}

func (n *notify) Wait() {
	n.wg.Wait()
}

func (n *notify) waitEvent() {
	for {
		select {
		case <-n.ctx.Done():
			return
		case subject := <-n.ch:
			for _, sender := range n.senders {
				tools.GoSafe(func() {
					defer n.wg.Done()
					err := sender.Send(subject)
					if err != nil {
						fmt.Println()
					}
				})
			}
		}
	}
}
