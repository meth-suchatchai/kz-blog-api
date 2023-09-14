package taximail

import (
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"log"
)

type Provide struct {
	ApiKey      string
	SecretKey   string
	URL         string
	SMSTemplate string
}

type defaultClient struct {
	Provide     *Provide
	SessionId   string
	SMSTemplate string
	Api         *resty.Client
}

func New(provide *Provide, api *resty.Client) (Client, error) {
	if provide.URL == "" && provide.ApiKey == "" && provide.SecretKey == "" {
		return nil, errors.New("provide not found")
	}

	client := defaultClient{
		Provide: provide,
		Api:     api,
	}

	err := client.Status()
	if err != nil {
		return nil, err
	}

	r, err := client.Login(LoginRequest{
		ApiKey:    provide.ApiKey,
		SecretKey: provide.SecretKey,
	})
	log.Print(r, err)
	client.SetSessionId(r.SessionId)
	client.SetSMSTemplate(provide.SMSTemplate)

	if client.SessionId == "" {
		return nil, errors.New("session not found or expired")
	}

	return &client, nil
}

func (c *defaultClient) SetSessionId(session string) {
	c.SessionId = session
}

func (c *defaultClient) SetSMSTemplate(templateId string) Client {
	c.SMSTemplate = templateId
	return c
}
