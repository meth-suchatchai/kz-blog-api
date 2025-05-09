package userservices

import (
	"github.com/gofiber/fiber/v2"
	usermodels "github.com/meth-suchatchai/kz-blog-api/app/user/models"
)

func (svc *defaultService) UserProfile(userId int64) (*usermodels.User, *fiber.Error) {
	return svc.userRepo.GetUser(userId)
}
