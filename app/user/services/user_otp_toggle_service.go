package userservices

import (
	"github.com/gofiber/fiber/v2"
	usermodels "github.com/meth-suchatchai/kz-blog-api/app/user/models"
)

func (svc *defaultService) UserOtpToggle(data *usermodels.User, enabled bool) (string, *fiber.Error) {
	sc := ""
	uri := ""
	if enabled {
		sc, uri = svc.totp.GenerateTOTP(data.Name)
	}

	err := svc.userRepo.UpdateTwoFactor(uint(data.Id), sc, enabled)
	if err != nil {
		return "", err
	}

	return uri, nil
}
