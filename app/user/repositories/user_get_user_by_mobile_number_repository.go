package userrepositories

import (
	"github.com/gofiber/fiber/v2"
	usermodels "github.com/meth-suchatchai/kz-blog-api/app/user/models"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
)

func (repo *defaultRepository) GetUserByMobileNumber(mobileNumber, countryCode string) (*usermodels.User, *fiber.Error) {
	user, err := repo.orm.GetUserByMobileNumber(mobileNumber)
	if err != nil {
		return nil, fiber.NewError(errors.ErrCodeInternalServer, err.Error())
	}

	newUser := usermodels.User{
		Id:           int64(user.ID),
		MobileNumber: user.MobileNumber,
		CountryCode:  user.CountryCode,
		Password:     user.PasswordEncrypted,
		Name:         user.FullName,
		Active:       user.IsActive,
		IsTFA:        user.TFEnable,
		TFACode:      user.TFCode,
	}

	return &newUser, nil
}
