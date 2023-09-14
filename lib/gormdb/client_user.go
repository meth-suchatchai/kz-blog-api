package gormdb

import (
	"github.com/gofiber/fiber/v2/log"
	dbmodels "github.com/kuroshibaz/lib/gormdb/models"
	"gorm.io/gorm"
)

func (c *DB) ListUser() ([]dbmodels.User, error) {
	var users []dbmodels.User
	tx := c.orm.Find(&users)
	if tx.Error != nil {
		return nil, nil
	}
	return users, nil
}

func (c *DB) GetUser(id uint) (*dbmodels.User, error) {
	var user dbmodels.User
	user.ID = id
	err := c.orm.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *DB) UpdateUser() bool {
	//c.orm.Updates()
	return false
}

func (c *DB) DeleteUser(id uint) bool {
	var user dbmodels.User
	user.ID = id
	err := c.orm.Delete(&user).Error
	if err != nil {
		log.Errorf("delete failed error: %d", id)
		return false
	}
	return true
}

func (c *DB) GetUserByMobileNumber(mobileNumber string) (*dbmodels.User, error) {
	var user dbmodels.User
	err := c.orm.Find(&user, "mobile_number = ?", mobileNumber).Error
	if err != nil {
		log.Errorf("get user failed: %v", err)
		return nil, err
	}

	return &user, nil
}

func (c *DB) CreateUser(data *dbmodels.User) (*dbmodels.User, error) {
	result := c.orm.Create(&data)
	if result.Error != nil {
		log.Errorf("create user failed: %v", result.Error)
		return nil, result.Error
	}

	return data, nil
}

func (c *DB) UpdateTFAColumn(enabled bool) error {
	var user dbmodels.User
	err := c.orm.Model(&user).Update("two_factor_enabled", enabled).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *DB) VerifyUser(id uint) error {
	var user = dbmodels.User{Model: gorm.Model{ID: id}}
	err := c.orm.Model(&user).Update("is_active", true).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *DB) GetUserPermission(userId uint, permissionId uint) bool {
	//log.Info("getUserPermission: ", userId, permissionId)
	var user = dbmodels.User{Model: gorm.Model{ID: userId}}
	err := c.orm.Preload("Permission").Find(&user).Error
	//log.Info(user)
	if err != nil {
		return false
	}

	roleExist := false
	for _, v := range user.Permission {
		if permissionId == v.ID {
			roleExist = true
		}
	}

	return roleExist
}
