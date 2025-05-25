package userservices

import (
	"github.com/gofiber/fiber/v2"
	usermodels "github.com/meth-suchatchai/kz-blog-api/app/user/models"
)

type Service interface {
	UserProfile(userId int64) (*usermodels.User, *fiber.Error)
	UserOtpToggle(data *usermodels.User, enabled bool) (string, *fiber.Error)
}
