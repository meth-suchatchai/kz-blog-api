package taximail

import (
	"github.com/gofiber/fiber/v2"
)

type Client interface {
	SetSessionId(session string)
	SendOTP(req OTPRequest) (*OTPResponse, *fiber.Error)
	VerifyOTP(req VerifyOTPRequest) (*VerifyOTPResponse, *fiber.Error)
	Status() *fiber.Error
	Login(req LoginRequest) (*LoginResponse, *fiber.Error)

	SetSMSTemplate(templateId string) Client
}
