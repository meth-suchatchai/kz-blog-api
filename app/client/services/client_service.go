package clientservices

import (
	"github.com/gofiber/fiber/v2"
	clientmodels "github.com/meth-suchatchai/kz-blog-api/app/client/models"
	usermodels "github.com/meth-suchatchai/kz-blog-api/app/user/models"
	kzjwt "github.com/meth-suchatchai/kz-blog-api/lib/jwt"
)

type Service interface {
	Register(data clientmodels.RegisterData) (*clientmodels.RegisterOTPUser, *fiber.Error)
	Login(data clientmodels.LoginData) (*usermodels.User, *kzjwt.AccessToken, *fiber.Error)
	VerifyOTP(data clientmodels.VerifyOTPData) *fiber.Error
	TwoFactorVerify(data clientmodels.TwoFactorVerifyData) *fiber.Error
}
