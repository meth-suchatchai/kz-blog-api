package gormdb

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/meth-suchatchai/kurostatemachine"
	dbmodels "github.com/meth-suchatchai/kz-blog-api/lib/gormdb/models"
)

func (c *defaultClient) CreateScene(data *dbmodels.Scene) (*dbmodels.Scene, error) {
	result := c.orm.Create(&data)
	if result.Error != nil {
		log.Errorf("create scene failed: %v", result.Error)
		return nil, result.Error
	}

	return data, nil
}

func (c *defaultClient) UpdateScene(id uint, params map[string]interface{}) error {
	return c.orm.Model(&dbmodels.Scene{}).Where("id = ?", id).Updates(params).Error
}

func (c *defaultClient) UpdateStatusScene(id uint, status kurostatemachine.State) error {
	return c.orm.Model(&dbmodels.Scene{}).Where("id = ?", id).Update("status", status).Error
}
