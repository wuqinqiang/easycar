// Package telegram is the telegram notification package.
package telegram

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
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
	req, err := http.NewRequest(http.MethodPost, api, nil)
	if err != nil {
		return err
	}
	req.Close = true
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Error response from Telegram - code [%d] - msg [%s]", resp.StatusCode, string(buf))
	}
	return nil
}
