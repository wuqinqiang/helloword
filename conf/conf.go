package conf

import (
	_ "embed"
	"io"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/wuqinqiang/helloword/notify"
	"github.com/wuqinqiang/helloword/notify/dingtalk"
	"github.com/wuqinqiang/helloword/notify/lark"
	"github.com/wuqinqiang/helloword/notify/telegram"
)

//go:embed conf.yml
var base []byte

func GetConf(filePath string) (*Settings, error) {
	var (
		settings Settings
		b        = base
	)

	if filePath != "" {
		file, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		defer file.Close() //nolint

		if b, err = io.ReadAll(file); err != nil {
			return nil, err
		}
	}

	err := yaml.Unmarshal(b, &settings)
	if err != nil {
		return nil, err
	}
	return &settings, nil
}

type Settings struct {
	GptToken string `yaml:"gptToken"`
	Notify
}

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
