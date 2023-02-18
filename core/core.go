package core

import (
	"github.com/wuqinqiang/helloword/generator"
	"github.com/wuqinqiang/helloword/notify"
)

type Core struct {
	generator generator.Generator
	notify    notify.Notify
}

func New(generator generator.Generator, notify notify.Notify) *Core {
	return &Core{
		generator: generator,
		notify:    notify,
	}
}
