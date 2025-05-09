package taximail

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
)

type OTPRequest struct {
	MobileName int    `json:"to"`
	SenderFrom string `json:"sms_template_key"`
}

type OTPResponse struct {
	Status string  `json:"status"`
	Code   int     `json:"code"`
	Data   OTPData `json:"data"`
	//ErrMsg string      `json:"err_msg,omitempty"`
}

type OTPData struct {
	MessageId        string  `json:"message_id"`
	RemainingBalance float64 `json:"remaining_balance"`
	MessagePrice     float64 `json:"message_price"`
	MessageCount     int     `json:"message_count"`
	OtpRefNo         string  `json:"otp_ref_no"`
}

func (c *defaultClient) SendOTP(req OTPRequest) (*OTPResponse, *fiber.Error) {
	req.SenderFrom = c.SMSTemplate
	res, err := c.Api.R().
		SetHeaders(map[string]string{
			"Authorization": "Bearer " + c.SessionId},
		).
		SetBody(req).
		Post(c.Provide.URL + "/v2/otp")
	if err != nil {
		return nil, errors.NewDefaultFiberError(err)
	}
	log.Info("SendOTP result: ", string(res.Body()))
	var result OTPResponse
	err = json.Unmarshal(res.Body(), &result)
	if err != nil {
		return nil, errors.NewDefaultFiberError(err)
	}

	if result.Code != 202 {
		return nil, errors.NewDefaultFiberMessageError("can't send OTP")
	}

	return &result, nil
}
