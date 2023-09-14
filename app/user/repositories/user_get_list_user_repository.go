package userrepositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kuroshibaz/lib/errors"
	dbmodels "github.com/kuroshibaz/lib/gormdb/models"
)

func (repo *defaultRepository) GetListUser() ([]dbmodels.User, *fiber.Error) {
	users, err := repo.cli.ListUser()
	if err != nil {
		return nil, fiber.NewError(errors.ErrCodeInternalServer, err.Error())
	}

	return users, nil
}
