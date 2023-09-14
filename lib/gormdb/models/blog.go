package dbmodels

import (
	"gorm.io/gorm"
	"time"
)

type Blog struct {
	gorm.Model
	Name             string    `gorm:"column:name;type:varchar(50)"`
	Content          string    `gorm:"column:content;type:text"`
	Slug             string    `gorm:"column:slug;type:varchar(255)"`
	Tag              []Tag     `gorm:"many2many:blog_tags"`
	Image            string    `gorm:"column:image_path"`
	PublishedAt      time.Time `gorm:"column:published_at"`
	Category         Category
	CategoryId       int
	SEO              SEO
	ShortDescription string `gorm:"column:short_description;type:varchar(128)"`
	Views            int
}
