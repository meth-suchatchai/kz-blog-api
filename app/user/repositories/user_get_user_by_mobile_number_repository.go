package userrepositories

import (
	"github.com/gofiber/fiber/v2"
	usermodels "github.com/kuroshibaz/app/user/models"
	"github.com/kuroshibaz/lib/errors"
)

func (repo *defaultRepository) GetUserByMobileNumber(mobileNumber, countryCode string) (*usermodels.User, *fiber.Error) {
	user, err := repo.cli.GetUserByMobileNumber(mobileNumber, countryCode)
	if err != nil {
		return nil, fiber.NewError(errors.ErrCodeInternalServer, err.Error())
	}

	newUser := usermodels.User{
		Id:           int64(user.ID),
		MobileNumber: user.MobileNumber,
		Password:     user.PasswordEncrypted,
		Name:         user.FullName,
		Active:       user.IsActive,
		IsTFA:        user.TFEnable,
		TFACode:      user.TFCode,
	}

	return &newUser, nil
}
