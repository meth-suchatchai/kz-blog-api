package userrepositories

import (
	"github.com/gofiber/fiber/v2"
	usermodels "github.com/meth-suchatchai/kz-blog-api/app/user/models"
	dbmodels "github.com/meth-suchatchai/kz-blog-api/lib/gormdb/models"
	kzjwt "github.com/meth-suchatchai/kz-blog-api/lib/jwt"
)

type Repository interface {
	CreateUser(data *dbmodels.User) (*usermodels.User, *fiber.Error)
	GetListUser() ([]dbmodels.User, *fiber.Error)
	GetUserByMobileNumber(mobileNumber, countryCode string) (*usermodels.User, *fiber.Error)
	UpdateTwoFactor(enabled bool) *fiber.Error
	VerifyUser(id int64) *fiber.Error
	GetUser(id int64) (*usermodels.User, *fiber.Error)

	GetUserAuthenticationByUserId(userId uint, token string) (*dbmodels.UserAuthentication, *fiber.Error)
	CreateOrUpdateUserAuthentication(user *usermodels.User, ac *kzjwt.AccessToken) *fiber.Error
}
