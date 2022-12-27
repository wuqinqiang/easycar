package lark

import (
	"encoding/json"
	"fmt"
	"time"

	. "github.com/wuqinqiang/easycar/tools"
)

// NotifyConfig is the lark notification configuration
type NotifyConfig struct {
	WebhookURL string `yaml:"webhook"`
}

// Send is the wrapper for SendLarkNotification
func (c NotifyConfig) Send(title, msg string) error {
	return c.SendLarkNotification(msg)
}

type b struct {
	MsgType string  `json:"msg_type"`
	Context Context `json:"content"`
}

type Context struct {
	Text string `json:"text"`
}

// SendLarkNotification will post to an 'Robot Webhook' url in Lark Apps. It accepts
// some text and the Lark robot will send it in group.
func (c NotifyConfig) SendLarkNotification(msg string) error {
	b := b{
		MsgType: "text",
	}
	b.Context.Text = msg

	resp, err := RestyCli.SetTimeout(5*time.Second).SetRetryCount(3).R().
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
