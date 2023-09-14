package dbmodels

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name        string `gorm:"column:name;type:varchar(50)"`
	Code        string `gorm:"column:code;type:varchar(20);uniqueIndex"`
	Description string `gorm:"column:description;type:text"`
	IsActive    bool   `gorm:"column:is_active;type:bool;default(true)"`
	Roles       []Role `gorm:"many2many:role_permissions"`
}
