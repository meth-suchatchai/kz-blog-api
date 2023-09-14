package dbmodels

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(50)"`
}
