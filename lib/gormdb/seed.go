package gormdb

import (
	"github.com/gofiber/fiber/v2/log"
	dbmodels "github.com/kuroshibaz/lib/gormdb/models"
)

func (c *defaultClient) Seed() {
	tx := c.orm.Create(&dbmodels.Role{
		Name:        "Admin",
		Description: "full access control",
		Permission: []dbmodels.Permission{
			{
				Name:        "Create Role",
				Code:        "CREATE_ROLE",
				Description: "",
				IsActive:    true,
			},
			{
				Name:        "Create Permission",
				Code:        "CREATE_PERMISSION",
				Description: "",
				IsActive:    true,
			},
		},
	})
	if tx.Error != nil {
		log.Error("seed error: ", tx.Error)
	}
}
