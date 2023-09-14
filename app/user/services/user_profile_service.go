package userservices

import (
	"github.com/gofiber/fiber/v2"
	usermodels "github.com/kuroshibaz/app/user/models"
)

func (svc *defaultService) UserProfile(userId int64) (*usermodels.User, *fiber.Error) {
	return svc.userRepo.GetUser(userId)
}
