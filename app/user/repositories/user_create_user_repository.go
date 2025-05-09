package userrepositories

import (
	"github.com/gofiber/fiber/v2"
	usermodels "github.com/meth-suchatchai/kz-blog-api/app/user/models"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
	dbmodels "github.com/meth-suchatchai/kz-blog-api/lib/gormdb/models"
)

func (repo *defaultRepository) CreateUser(data *dbmodels.User) (*usermodels.User, *fiber.Error) {
	user, err := repo.orm.CreateUser(data)
	if err != nil {
		return nil, fiber.NewError(errors.ErrCodeInternalServer, err.Error())
	}

	u := &usermodels.User{
		Id:           int64(data.ID),
		MobileNumber: user.MobileNumber,
		Password:     user.PasswordEncrypted,
		Name:         user.FullName,
		Active:       user.IsActive,
		IsTFA:        user.TFEnable,
		TFACode:      user.TFCode,
	}

	return u, nil
}
