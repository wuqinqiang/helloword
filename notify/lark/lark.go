package lark

import (
	"encoding/json"
	"fmt"
	"github.com/wuqinqiang/helloword/notify/base"
	"time"

	. "github.com/wuqinqiang/helloword/tools"
)

// NotifyConfig is the lark notification configuration
type NotifyConfig struct {
	WebhookURL string `yaml:"webhook"`
}

// Send is the wrapper for SendLarkNotification
func (c NotifyConfig) Send(subject base.Subject) error {
	return c.SendLarkNotification(subject)
}

// SendLarkNotification will post to an 'Robot Webhook' url in Lark Apps. It accepts
// some text and the Lark robot will send it in group.
func (c NotifyConfig) SendLarkNotification(subject base.Subject) error {
	b := body{
		MsgType: "text",
	}
	b.Context.Text = subject.Text()

	resp, err := Resty.SetTimeout(5*time.Second).SetRetryCount(3).R().
		SetHeader("Content-Type", "application/json").
		SetBody(b).Post(c.WebhookURL)
	if err != nil {
		return err
	}

	ret := make(map[string]interface{})
	err = json.Unmarshal(resp.Body(), &ret)
	if err != nil {
		return fmt.Errorf("error response from Lark [%d] - [%s]", resp.StatusCode(), string(resp.Body()))
	}
	// Server returns {"Extra":null,"StatusCode":0,"StatusMessage":"success"} on success
	// otherwise it returns {"code":9499,"msg":"Bad Request","data":{}}
	if statusCode, ok := ret["StatusCode"].(float64); !ok || statusCode != 0 {
		code, _ := ret["code"].(float64)
		msg, _ := ret["msg"].(string)
		return fmt.Errorf("error response from Lark - code [%d] - msg [%v]", int(code), msg)
	}
	return nil
}
