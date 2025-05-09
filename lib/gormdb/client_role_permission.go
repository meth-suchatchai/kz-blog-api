package gormdb

import (
	"github.com/gofiber/fiber/v2/log"
	dbmodels "github.com/meth-suchatchai/kz-blog-api/lib/gormdb/models"
	"gorm.io/gorm"
)

func (c *defaultClient) CreateRole(r *dbmodels.Role) error {
	return c.orm.Create(r).Error
}

func (c *defaultClient) CreatePermission(r *dbmodels.Permission) error {
	return c.orm.Create(r).Error
}

func (c *defaultClient) GetRoles(opts ...int) (*[]dbmodels.Role, error) {
	var roles []dbmodels.Role
	page := 1
	limit := 100
	offset := 0
	if len(opts) <= 0 && len(opts) != 2 {
		page = opts[0]
		limit = opts[1]
		offset = limit * (page - 1)
	}

	err := c.orm.Find(&roles).Limit(limit).Offset(offset).Error
	if err != nil {
		return nil, err
	}

	return &roles, nil
}

func (c *defaultClient) GetRolePermission() (*[]dbmodels.Role, error) {
	var roles []dbmodels.Role
	err := c.orm.Preload("Permission", "is_active = ?", true).
		Find(&roles).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}

	return &roles, nil
}

func (c *defaultClient) GetPermission(permissionCode string) (*dbmodels.Permission, error) {
	var pem dbmodels.Permission
	err := c.orm.First(&pem, "code = ?", permissionCode).Error
	if err != nil {
		return nil, err
	}

	return &pem, nil
}

func (c *defaultClient) AssignRoleToUser(roleId uint, userId uint) (*dbmodels.Role, error) {
	role := dbmodels.Role{Model: gorm.Model{ID: roleId}}
	tx := c.orm.Begin()
	tx.Preload("Permission", "is_active = ?", true).
		First(&role)

	user := dbmodels.User{Model: gorm.Model{ID: userId}}
	if tx.Error != nil {
		return nil, tx.Error
	}

	tx.FirstOrCreate(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	user.Permission = append(role.Permission)
	tx.Save(&user)

	if tx.Error != nil {
		return nil, tx.Error
	}
	defer tx.Commit()
	//c.orm.Create()
	log.Info(role)
	log.Info(user)
	return nil, nil
}

func (c *defaultClient) AssignPermissionToRole(role string) {
	switch role {
	case "Admin":
		permissions, err := c.allPermissionGet()
		if err != nil {
			log.Error("can't get all permissions")
		}

		r := dbmodels.Role{}
		tx := c.orm.Begin()
		tx.First(&r, "name = ?", role)
		if tx.Error != nil {
			log.Error("error query: ", tx.Error)
		}

		r.Permission = append(permissions)
		tx.Save(&r)

		defer tx.Commit()
	default:
		log.Info("no role found")
	}
}

func (c *defaultClient) allPermissionGet() ([]dbmodels.Permission, error) {
	var pems []dbmodels.Permission
	err := c.orm.Find(&pems, "is_active = true").Error
	return pems, err
}
