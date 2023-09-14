package gormdb

import (
	"github.com/gofiber/fiber/v2/log"
	dbmodels "github.com/kuroshibaz/lib/gormdb/models"
	"gorm.io/gorm"
	"time"
)

func (c *DB) GetBlogById(id uint) (*dbmodels.Blog, error) {
	var blog = dbmodels.Blog{Model: gorm.Model{ID: id}}
	err := c.orm.First(&blog).Error
	if err != nil {
		return nil, err
	}

	return &blog, nil
}

func (c *DB) GetContentBySlug(slug string) (*dbmodels.Blog, error) {
	var blog dbmodels.Blog
	err := c.orm.First(&blog, "slug = ?", slug).Error
	if err != nil {
		return nil, err
	}

	return &blog, nil
}

func (c *DB) CreateBlog(data *dbmodels.Blog) error {
	data.PublishedAt = time.Now()
	data.Views = 0
	data.ShortDescription = ""
	return c.orm.Create(data).Error
}

func (c *DB) DeleteBlog(id uint) error {
	var blog = dbmodels.Blog{Model: gorm.Model{ID: id}}
	return c.orm.Delete(blog).Error
}

func (c *DB) UpdateBlog(data *dbmodels.Blog) error {
	err := c.orm.Save(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *DB) ListBlog(opts ...int) (*[]dbmodels.Blog, error) {
	var blogs []dbmodels.Blog
	page := 1
	limit := 100
	offset := 0
	if len(opts) <= 0 && len(opts) != 2 {
		page = opts[0]
		limit = opts[1]
		offset = limit * (page - 1)
	}

	err := c.orm.Find(&blogs).Limit(limit).Offset(offset).Error
	if err != nil {
		return nil, err
	}

	log.Info(blogs)
	return &blogs, nil
}

func (c *DB) ListPopularTag() (*[]dbmodels.Tag, error) {
	var tags []dbmodels.Tag
	err := c.orm.Order("ord desc").Find(&tags).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &[]dbmodels.Tag{}, nil
		}
		return nil, err
	}

	return &tags, nil
}

func (c *DB) CreateTag(data *dbmodels.Tag) error {
	return c.orm.Create(data).Error
}

func (c *DB) CreateCategory(data *dbmodels.Category) error {
	return c.orm.Create(data).Error
}

func (c *DB) ListCategory() (*[]dbmodels.Category, error) {
	var categories []dbmodels.Category
	err := c.orm.Find(&categories).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &[]dbmodels.Category{}, nil
		}
		return nil, err
	}
	return &categories, nil
}

func (c *DB) UpdateCategory(data *dbmodels.Category) error {
	return c.orm.Save(&data).Error
}

func (c *DB) CountViews(slug string) (int, error) {
	var blog dbmodels.Blog

	tx := c.orm.First(&blog, "slug = ?", slug)
	blog.Views += 1

	err := tx.Updates(&blog).Error
	if err != nil {
		return 0, err
	}

	return blog.Views, nil
}
