package userrepositories

import (
	"github.com/gofiber/fiber/v2"
	usermodels "github.com/kuroshibaz/app/user/models"
	dbmodels "github.com/kuroshibaz/lib/gormdb/models"
)

type Repository interface {
	CreateUser(data *dbmodels.User) (*usermodels.User, *fiber.Error)
	GetListUser() ([]dbmodels.User, *fiber.Error)
	GetUserByMobileNumber(mobileNumber string) (*usermodels.User, *fiber.Error)
	UpdateTwoFactor(enabled bool) *fiber.Error
	VerifyUser(id int64) *fiber.Error
	GetUser(id int64) (*usermodels.User, *fiber.Error)
}
