package dingtalk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/wuqinqiang/easycar/logging"
)

// NotifyConfig is the dingtalk notification configuration
type NotifyConfig struct {
	WebhookURL string `yaml:"webhook"`
	SignSecret string `yaml:"secret,omitempty"`
}

// Send will post to an 'Robot Webhook' url in Dingtalk Apps. It accepts
// some text and the Dingtalk robot will send it in group.
func (c NotifyConfig) Send(title, msg string) error {
	title = "**" + title + "**"
	// It will be better to escape the msg.
	msgContent := fmt.Sprintf(`
	{
		"msgtype": "markdown",
		"markdown": {
			"title": "%s",
			"text": "%s"
		}
	}
	`, title, msg)
	req, err := http.NewRequest(http.MethodPost, c.addSign(c.WebhookURL, c.SignSecret), bytes.NewBuffer([]byte(msgContent)))
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
	if err != nil || ret["errmsg"] != "ok" {
		return fmt.Errorf("error response from Dingtalk [%d] - [%s]", resp.StatusCode, string(buf))
	}
	return nil
}

// add sign for url by secret
func (c NotifyConfig) addSign(webhookURL string, secret string) string {
	webhook := webhookURL
	if secret != "" {
		timestamp := time.Now().UnixMilli()
		stringToSign := fmt.Sprint(timestamp, "\n", secret)
		h := hmac.New(sha256.New, []byte(secret))
		h.Write([]byte(stringToSign))
		sign := url.QueryEscape(base64.StdEncoding.EncodeToString(h.Sum(nil)))
		webhook = fmt.Sprint(webhookURL, "&timestamp=", timestamp, "&sign="+sign)
	}
	logging.Debugf("Dingtalk webhook: %s", webhook)
	return webhook
}
