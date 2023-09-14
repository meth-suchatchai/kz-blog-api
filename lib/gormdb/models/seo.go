package dbmodels

import "gorm.io/gorm"

type SEO struct {
	gorm.Model
	BlogId          uint   `gorm:"uniqueIndex"`
	MetaTitle       string `gorm:"column:meta_title;type:varchar(255)"`
	MetaDescription string `gorm:"column:meta_description;type:text"`
}
