package userrepositories

import (
	"github.com/gofiber/fiber/v2"
	usermodels "github.com/kuroshibaz/app/user/models"
	dbmodels "github.com/kuroshibaz/lib/gormdb/models"
	kzjwt "github.com/kuroshibaz/lib/jwt"
	"gorm.io/gorm"
)

func (repo *defaultRepository) CreateOrUpdateUserAuthentication(user *usermodels.User, ac *kzjwt.AccessToken) *fiber.Error {
	err := repo.orm.CreateOrUpdateUserAuthentication(&dbmodels.User{
		Model: gorm.Model{
			ID: uint(user.Id),
		},
		MobileNumber: user.MobileNumber,
		CountryCode:  user.CountryCode,
		FullName:     user.Name,
	}, &dbmodels.UpdateUserAuthentication{
		UserId:             uint(user.Id),
		MobileNumber:       user.MobileNumber,
		CountryCode:        user.CountryCode,
		AccessToken:        &ac.AccessToken,
		AccessTokenExpire:  &ac.AccessTokenExpire,
		RefreshToken:       &ac.RefreshToken,
		RefreshTokenExpire: &ac.RefreshTokenExpire,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
