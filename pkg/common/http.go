package common

import (
	"time"

	"github.com/go-resty/resty/v2"
)

var RestyClient = httpClient{}

func init() {
	client := resty.New()
	RestyClient.Client = client
	//RestyClient.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
	//	return nil
	//})
	//
	//RestyClient.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
	//	return nil
	//})
}

type httpClient struct {
	*resty.Client
}

type HttpOption func(client *httpClient)

func SetTimeOut(time time.Duration) HttpOption {
	return func(client *httpClient) {
		client.SetTimeout(time)
	}
}

func (h *httpClient) PostJson(uri string, body interface{}, result interface{}, options ...HttpOption) error {
	for _, option := range options {
		option(h)
	}
	_, err := h.R().SetHeader("Context-Type", "application/json").
		SetBody(body).SetResult(&result).Post(uri)
	return err
}
