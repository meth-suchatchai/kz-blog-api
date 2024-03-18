package userrepositories

import (
	"github.com/gofiber/fiber/v2"
	dbmodels "github.com/kuroshibaz/lib/gormdb/models"
)

func (repo *defaultRepository) GetUserAuthenticationByUserId(userId uint, token string) (*dbmodels.UserAuthentication, *fiber.Error) {
	auth, err := repo.orm.GetUserAuthenticationById(userId, token)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return auth, nil
}
