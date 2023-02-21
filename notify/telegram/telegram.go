package telegram

import (
	"fmt"
	"net/url"
	"time"

	"github.com/wuqinqiang/helloword/notify/base"

	. "github.com/wuqinqiang/helloword/tools"
)

// NotifyConfig is the telegram notification configuration
type NotifyConfig struct {
	Token  string `yaml:"token"`
	ChatID string `yaml:"chat_id"`
}

// Send is the wrapper for SendTelegramNotification
func (c NotifyConfig) Send(subject base.Subject) error {
	return c.SendTelegramNotification(subject)
}

// SendTelegramNotification will send the notification to telegram.
func (c NotifyConfig) SendTelegramNotification(subject base.Subject) error {
	api := "https://api.telegram.org/bot" + c.Token +
		"/sendMessage?&chat_id=" + c.ChatID +
		"&parse_mode=markdown" +
		"&text=" + url.QueryEscape(subject.Text())

	resp, err := Resty.SetTimeout(5*time.Second).SetRetryCount(3).R().
		SetHeader("Content-Type", "application/json").
		Post(api)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("error response from Telegram - code [%d]", resp.StatusCode())
	}
	return nil
}
