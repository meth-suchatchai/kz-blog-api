package gormdb

import (
	"github.com/gofiber/fiber/v2/log"
	dbmodels "github.com/kuroshibaz/lib/gormdb/models"
	"gorm.io/gorm"
	"time"
)

func (c *defaultClient) GetUserAuthenticationByMobile(mobileNumber, countryCode string) (*dbmodels.UserAuthentication, error) {
	var userAuth = &dbmodels.UserAuthentication{}
	log.Infof("getUserAuthentication data: %v, %v", mobileNumber, countryCode)
	exec := c.orm.Where("mobile_number = ? AND country_code = ?", mobileNumber, countryCode).First(&userAuth)
	if exec.Error != nil {
		log.Errorf("get user failed: %v", exec.Error)
		return nil, exec.Error
	}

	return userAuth, nil
}

func (c *defaultClient) GetUserAuthenticationById(userId uint, token string) (*dbmodels.UserAuthentication, error) {
	var userAuth = &dbmodels.UserAuthentication{}
	log.Infof("getUserAuthentication data: %v token: %v", userId, token)
	exec := c.orm.Where("user_id = ? and access_token = ? and access_token_expire >= ?", userId, token, time.Now()).First(&userAuth)
	if exec.Error != nil {
		log.Errorf("get user failed: %v", exec.Error)
		return nil, exec.Error
	}

	return userAuth, nil
}

func (c *defaultClient) CreateUserAuthentication(data *dbmodels.UserAuthentication) (*dbmodels.UserAuthentication, error) {
	result := c.orm.Create(&data)
	if result.Error != nil {
		log.Errorf("create user authentication: %v", result.Error)
		return nil, result.Error
	}

	return data, nil
}

func (c *defaultClient) UpdateUserAuthentication(field *dbmodels.UpdateUserAuthentication) bool {
	var user dbmodels.UserAuthentication
	updates := map[string]interface{}{}

	if field.AccessToken != nil {
		updates["access_token"] = *field.AccessToken
	}

	if field.AccessTokenExpire != nil {
		updates["access_token_expire"] = *field.AccessTokenExpire
	}

	if field.RefreshToken != nil {
		updates["refresh_token"] = *field.RefreshToken
	}

	if field.RefreshTokenExpire != nil {
		updates["refresh_token_expire"] = *field.RefreshTokenExpire
	}

	exec := c.orm.Model(&user).Updates(updates)
	if exec.Error != nil {
		log.Errorf("update user authentication failed: %v", exec.Error)
		return false
	}

	return true
}

func (c *defaultClient) CreateOrUpdateUserAuthentication(user *dbmodels.User, auth *dbmodels.UpdateUserAuthentication) error {
	tx := c.orm.Begin()
	defer tx.Commit()
	var m dbmodels.UserAuthentication

	exec := tx.First(&m, "user_id = ?", user.ID)
	log.Info("exec: ", exec)
	if exec.Error != nil {
		if exec.Error == gorm.ErrRecordNotFound {
			//Create it
			m = dbmodels.UserAuthentication{
				UserId:             user.ID,
				MobileNumber:       user.MobileNumber,
				CountryCode:        user.CountryCode,
				AccessToken:        *auth.AccessToken,
				AccessTokenExpire:  *auth.AccessTokenExpire,
				RefreshToken:       *auth.RefreshToken,
				RefreshTokenExpire: *auth.RefreshTokenExpire,
			}
			tx.Create(&m)
			if tx.Error != nil {
				log.Errorf("CreateOrUpdateUserAuthentication create user authentication failed: %v", tx.Error)
				tx.Rollback()
				return tx.Error
			}
			return nil
		} else {
			log.Errorf("CreateOrUpdateUserAuthentication error: %v", tx.Error)
			tx.Rollback()
			return tx.Error
		}
	} else {
		updates := map[string]interface{}{}

		if auth.AccessToken != nil {
			updates["access_token"] = *auth.AccessToken
		}

		if auth.AccessTokenExpire != nil {
			updates["access_token_expire"] = *auth.AccessTokenExpire
		}

		if auth.RefreshToken != nil {
			updates["refresh_token"] = *auth.RefreshToken
		}

		if auth.RefreshTokenExpire != nil {
			updates["refresh_token_expire"] = *auth.RefreshTokenExpire
		}

		tx.Model(&m).Where("user_id = ?", user.ID).Updates(updates)
		if tx.Error != nil {
			log.Errorf("CreateOrUpdateUserAuthentication update user authentication failed: %v", tx.Error)
			tx.Rollback()
			return tx.Error
		}

		return nil
	}
}
