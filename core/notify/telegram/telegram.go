// Package telegram is the telegram notification package.
package telegram

import (
	"fmt"
	"net/url"
	"time"

	. "github.com/wuqinqiang/easycar/tools"
)

// NotifyConfig is the telegram notification configuration
type NotifyConfig struct {
	Token  string `yaml:"token"`
	ChatID string `yaml:"chat_id"`
}

// Send is the wrapper for SendTelegramNotification
func (c NotifyConfig) Send(title, text string) error {
	return c.SendTelegramNotification(text)
}

// SendTelegramNotification will send the notification to telegram.
func (c NotifyConfig) SendTelegramNotification(text string) error {
	api := "https://api.telegram.org/bot" + c.Token +
		"/sendMessage?&chat_id=" + c.ChatID +
		"&parse_mode=markdown" +
		"&text=" + url.QueryEscape(text)

	resp, err := RestyCli.SetTimeout(5*time.Second).SetRetryCount(3).R().
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
