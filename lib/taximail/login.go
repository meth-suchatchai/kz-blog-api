package taximail

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
)

type LoginRequest struct {
	ApiKey    string `json:"api_key"`
	SecretKey string `json:"secret_key"`
}

type LoginResponse struct {
	Status    string `json:"status"`
	Code      int    `json:"code"`
	LoginData `json:"data"`
}

type LoginData struct {
	Expire    int64  `json:"expire"`
	UserType  string `json:"user_type"`
	SessionId string `json:"session_id"`
}

func (c *defaultClient) Login(req LoginRequest) (*LoginResponse, *fiber.Error) {
	res, err := c.Api.R().SetBody(req).Post(c.Provide.URL + "/v2/user/login")
	if err != nil {
		return nil, errors.NewDefaultFiberError(err)
	}
	var result LoginResponse
	err = json.Unmarshal(res.Body(), &result)
	if err != nil {
		return nil, errors.NewDefaultFiberError(err)
	}

	if result.Code != 201 {
		return nil, errors.NewDefaultFiberMessageError("can't not login")
	}

	return &result, nil
}
