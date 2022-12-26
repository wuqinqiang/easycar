package lark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// NotifyConfig is the lark notification configuration
type NotifyConfig struct {
	WebhookURL string `yaml:"webhook"`
}

// Send is the wrapper for SendLarkNotification
func (c NotifyConfig) Send(title, msg string) error {
	return c.SendLarkNotification(msg)
}

// SendLarkNotification will post to an 'Robot Webhook' url in Lark Apps. It accepts
// some text and the Lark robot will send it in group.
func (c NotifyConfig) SendLarkNotification(msg string) error {
	req, err := http.NewRequest(http.MethodPost, c.WebhookURL, bytes.NewBuffer([]byte(msg)))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Close = true

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	ret := make(map[string]interface{})
	err = json.Unmarshal(buf, &ret)
	if err != nil {
		return fmt.Errorf("error response from Lark [%d] - [%s]", resp.StatusCode, string(buf))
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
