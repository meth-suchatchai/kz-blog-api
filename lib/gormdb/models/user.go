package dbmodels

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	MobileNumber      string       `gorm:"column:mobile_number;type:varchar(10);uniqueIndex"`
	CountryCode       string       `gorm:"column:country_code;type:varchar(5);not null"`
	FullName          string       `gorm:"column:full_name;type:varchar(255)"`
	IsActive          bool         `gorm:"column:is_active;type:bool;default(false)"`
	PasswordEncrypted string       `gorm:"column:password_encrypted;type:varchar(32);not null"`
	TFEnable          bool         `gorm:"column:two_factor_enabled;type:bool;default(false)"`
	TFCode            string       `gorm:"column:two_factor_code;type:varchar(32);default:''"`
	Permission        []Permission `gorm:"many2many:user_permissions"`
}
