package dbmodels

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string       `gorm:"column:name;type:varchar(50)"`
	Description string       `gorm:"column:description;type:text"`
	Permission  []Permission `gorm:"many2many:role_permissions"`
}
