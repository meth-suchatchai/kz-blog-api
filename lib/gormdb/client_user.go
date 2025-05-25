package gormdb

import (
	"github.com/gofiber/fiber/v2/log"
	dbmodels "github.com/meth-suchatchai/kz-blog-api/lib/gormdb/models"
	"gorm.io/gorm"
)

func (c *defaultClient) ListUser() ([]dbmodels.User, error) {
	var users []dbmodels.User
	tx := c.orm.Find(&users)
	if tx.Error != nil {
		return nil, nil
	}
	return users, nil
}

func (c *defaultClient) GetUser(id uint) (*dbmodels.User, error) {
	var user dbmodels.User
	user.ID = id
	err := c.orm.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *defaultClient) UpdateUser(id uint, params map[string]interface{}) error {
	return c.orm.Model(&dbmodels.User{}).Where("id = ?", id).Updates(params).Error
}

func (c *defaultClient) DeleteUser(id uint) bool {
	var user dbmodels.User
	user.ID = id
	err := c.orm.Delete(&user).Error
	if err != nil {
		log.Errorf("delete failed error: %d", id)
		return false
	}
	return true
}

func (c *defaultClient) GetUserByMobileNumber(mobileNumber string) (*dbmodels.User, error) {
	var user = &dbmodels.User{}
	exec := c.orm.Where("mobile_number = ?", mobileNumber).First(&user)
	if exec.Error != nil {
		log.Errorf("get user failed: %v", exec.Error)
		return nil, exec.Error
	}

	return user, nil
}

func (c *defaultClient) CreateUser(data *dbmodels.User) (*dbmodels.User, error) {
	result := c.orm.Create(&data)
	if result.Error != nil {
		log.Errorf("create user failed: %v", result.Error)
		return nil, result.Error
	}

	return data, nil
}

func (c *defaultClient) UpdateTFAColumn(id uint, secretKey string, enabled bool) error {
	var user = dbmodels.User{Model: gorm.Model{ID: id}}
	err := c.orm.Model(&user).Updates(dbmodels.User{TFEnable: enabled, TFCode: secretKey}).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *defaultClient) VerifyUser(id uint) error {
	var user = dbmodels.User{Model: gorm.Model{ID: id}}
	err := c.orm.Model(&user).Update("is_active", true).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *defaultClient) GetUserPermission(userId uint, permissionId uint) bool {
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
