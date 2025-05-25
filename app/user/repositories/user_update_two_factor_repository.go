package userrepositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
	"log"
)

func (repo *defaultRepository) UpdateTwoFactor(id uint, secretKey string, enabled bool) *fiber.Error {
	log.Print("UserOtpToggle: ", id, secretKey, enabled)
	err := repo.orm.UpdateTFAColumn(id, secretKey, enabled)
	log.Println("UpdateTFAColumn: ", err)
	if err != nil {
		return fiber.NewError(errors.ErrCodeInternalServer, err.Error())
	}

	return nil
}
