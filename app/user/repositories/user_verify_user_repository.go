package userrepositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
)

func (repo *defaultRepository) VerifyUser(id int64) *fiber.Error {
	err := repo.orm.VerifyUser(uint(id))
	if err != nil {
		return fiber.NewError(errors.ErrCodeInternalServer, err.Error())
	}

	return nil
}
