package taximail

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/kuroshibaz/lib/errors"
)

type VerifyOTPRequest struct {
	MessageId    string `json:"message_id"`
	OTPCode      string `json:"otp_code"`
	MobileNumber int    `json:"to"`
}

type VerifyOTPResponse struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	ErrMsg string      `json:"err_msg"`
}

func (c *defaultClient) VerifyOTP(req VerifyOTPRequest) (*VerifyOTPResponse, *fiber.Error) {
	res, err := c.Api.R().
		SetHeader("Authorization", "Bearer "+c.SessionId).
		SetQueryParams(map[string]string{
			"to":         fmt.Sprintf("%v", req.MobileNumber),
			"otp_code":   req.OTPCode,
			"message_id": req.MessageId,
		}).
		Get(c.Provide.URL + fmt.Sprintf("/v2/otp/verify/%v", req.MessageId))
	log.Info("VerifyOTP error: ", string(res.Body()), err)
	if err != nil {
		return nil, errors.NewDefaultFiberError(err)
	}

	var result VerifyOTPResponse
	err = json.Unmarshal(res.Body(), &result)
	if err != nil {
		return nil, errors.NewDefaultFiberError(err)
	}

	if result.Code != 202 {
		return nil, errors.NewDefaultFiberMessageError(result.ErrMsg)
	}

	return &result, nil
}
