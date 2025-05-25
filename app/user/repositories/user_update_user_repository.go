package userrepositories

import (
	"github.com/gofiber/fiber/v2"
	usermodels "github.com/meth-suchatchai/kz-blog-api/app/user/models"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
)

func (repo *defaultRepository) UpdateUser(id uint, data *usermodels.UpdateUser) *fiber.Error {
	params := make(map[string]interface{})
	if data.Name != nil {
		params["name"] = *data.Name
	}

	if data.CountryCode != nil {
		params["country_code"] = *data.CountryCode
	}

	if data.MobileNumber != nil {
		params["mobile_number"] = *data.MobileNumber
	}

	err := repo.orm.UpdateUser(id, params)
	if err != nil {
		return fiber.NewError(errors.ErrCodeInternalServer, err.Error())
	}

	return nil
}
