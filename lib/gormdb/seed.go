package gormdb

import (
	"github.com/gofiber/fiber/v2/log"
	dbmodels "github.com/kuroshibaz/lib/gormdb/models"
	"github.com/kuroshibaz/lib/utils"
)

var createPermissions = []dbmodels.Permission{
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
	{
		Name:        "Read Role Permission",
		Code:        "READ_ROLE_PERMISSION",
		Description: "",
		IsActive:    true,
	},
	{
		Name:        "Admin Read Blog",
		Code:        "READ_BLOG",
		Description: "",
		IsActive:    true,
	},
	{
		Name:        "Admin Create Blog",
		Code:        "CREATE_BLOG",
		Description: "",
		IsActive:    true,
	},
	{
		Name:        "Admin Update Blog",
		Code:        "UPDATE_BLOG",
		Description: "",
		IsActive:    true,
	},
	{
		Name:        "Admin Delete Blog",
		Code:        "DELETE_BLOG",
		Description: "",
		IsActive:    true,
	},
	{
		Name:        "Admin Audit",
		Code:        "AUDIT",
		Description: "",
		IsActive:    true,
	},
}

func (c *defaultClient) Seed() {
	tx := c.orm.Create(createPermissions)

	tx = c.orm.Create(&dbmodels.Role{
		Name:        "Admin",
		Description: "full access control",
		Permission:  createPermissions,
	})
	if tx.Error != nil {
		log.Error("kzcli error: ", tx.Error)
	}

	encryptdPass := utils.EncryptedHash("qwertyuiop")
	tx = c.orm.Create(&dbmodels.User{
		MobileNumber:      "8023736019",
		CountryCode:       "81",
		FullName:          "Kuroshibz",
		IsActive:          true,
		PasswordEncrypted: encryptdPass,
		TFEnable:          false,
		TFCode:            "",
		Permission:        createPermissions,
	})
}
