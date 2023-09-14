package gormdb

import dbmodels "github.com/kuroshibaz/lib/gormdb/models"

func (c *DB) Migrate() error {
	err := c.orm.AutoMigrate(&dbmodels.User{})
	err = c.orm.AutoMigrate(&dbmodels.Permission{}, &dbmodels.Role{})
	err = c.orm.AutoMigrate(&dbmodels.Blog{}, &dbmodels.Tag{}, &dbmodels.Category{}, &dbmodels.SEO{})
	if err != nil {
		return err
	}

	return nil
}
