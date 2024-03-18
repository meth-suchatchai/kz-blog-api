package userrepositories

import (
	"github.com/gofiber/fiber/v2"
	usermodels "github.com/kuroshibaz/app/user/models"
	"github.com/kuroshibaz/lib/errors"
)

func (repo *defaultRepository) GetUser(id int64) (*usermodels.User, *fiber.Error) {
	user, err := repo.orm.GetUser(uint(id))
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

	//err = mapstructure.Decode(&user, &newUser)
	//if err != nil {
	//	return nil, fiber.NewError(errors.ErrCodeInternalServer, err.Error())
	//}

	return &newUser, nil
}
