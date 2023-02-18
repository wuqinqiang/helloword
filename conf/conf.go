package conf

import (
	"github.com/wuqinqiang/helloword/notify"
	"github.com/wuqinqiang/helloword/notify/dingtalk"
	"github.com/wuqinqiang/helloword/notify/lark"
	"github.com/wuqinqiang/helloword/notify/telegram"
)

// Notify Config
type Notify struct {
	Lark     lark.NotifyConfig     `yaml:"lark"`
	Tg       telegram.NotifyConfig `yaml:"tg"`
	Dingtalk dingtalk.NotifyConfig `yaml:"dingtalk"`
}

func (n *Notify) Senders() (senders []notify.Sender) {
	if n.Tg.Token != "" && n.Tg.ChatID != "" {
		senders = append(senders, n.Tg)
	}
	if n.Lark.WebhookURL != "" {
		senders = append(senders, n.Lark)
	}
	if n.Dingtalk.SignSecret != "" && n.Dingtalk.WebhookURL != "" {
		senders = append(senders, n.Dingtalk)
	}
	return
}
