package kzline

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type defaultClient struct {
	LineApi     string
	BotApi      string
	AccessToken string
	restyClient *resty.Client
}

func NewLineNotification(botUrl, apiUrl, token string, restyClient *resty.Client) LineNotification {
	return &defaultClient{BotApi: botUrl, LineApi: apiUrl, AccessToken: token, restyClient: restyClient}
}

func (c *defaultClient) getHeaderToken() string {
	return fmt.Sprintf("Bearer %s", c.AccessToken)
}
