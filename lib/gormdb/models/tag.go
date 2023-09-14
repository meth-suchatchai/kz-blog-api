package dbmodels

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name  string `gorm:"column:name;type:varchar(50)"`
	Ord   int    `gorm:"column:ord;"`
	Blogs []Blog `gorm:"many2many:blog_tags"`
}
