package gormdb

import dbmodels "github.com/kuroshibaz/lib/gormdb/models"

func (c *defaultClient) Migrate() error {
	err := c.orm.AutoMigrate(&dbmodels.User{})
	err = c.orm.AutoMigrate(&dbmodels.Permission{}, &dbmodels.Role{})
	err = c.orm.AutoMigrate(&dbmodels.Blog{}, &dbmodels.Tag{}, &dbmodels.Category{}, &dbmodels.SEO{})
	err = c.orm.AutoMigrate(&dbmodels.UserAuthentication{})
	if err != nil {
		return err
	}

	return nil
}
