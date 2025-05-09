package userrepositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
	dbmodels "github.com/meth-suchatchai/kz-blog-api/lib/gormdb/models"
)

func (repo *defaultRepository) GetListUser() ([]dbmodels.User, *fiber.Error) {
	users, err := repo.orm.ListUser()
	if err != nil {
		return nil, fiber.NewError(errors.ErrCodeInternalServer, err.Error())
	}

	return users, nil
}
