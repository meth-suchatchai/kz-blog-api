package userservices

import (
	"github.com/gofiber/fiber/v2"
	usermodels "github.com/kuroshibaz/app/user/models"
)

type Service interface {
	UserProfile(userId int64) (*usermodels.User, *fiber.Error)
}
