package clientservices

import (
	"github.com/gofiber/fiber/v2"
	clientmodels "github.com/kuroshibaz/app/client/models"
	usermodels "github.com/kuroshibaz/app/user/models"
	kzjwt "github.com/kuroshibaz/lib/jwt"
)

type Service interface {
	Register(data clientmodels.RegisterData) (*clientmodels.RegisterOTPUser, *fiber.Error)
	Login(data clientmodels.LoginData) (*usermodels.User, *kzjwt.AccessToken, *fiber.Error)
	VerifyOTP(data clientmodels.VerifyOTPData) *fiber.Error
}
