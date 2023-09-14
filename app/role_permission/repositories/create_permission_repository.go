package rprepositories

import (
	"github.com/gofiber/fiber/v2"
	rpmodels "github.com/kuroshibaz/app/role_permission/models"
	"github.com/kuroshibaz/lib/errors"
	dbmodels "github.com/kuroshibaz/lib/gormdb/models"
)

func (repo *defaultRepository) CreatePermission(data *rpmodels.Permission) *fiber.Error {
	err := repo.db.CreatePermission(&dbmodels.Permission{
		Name:        data.Name,
		Code:        data.Code,
		Description: data.Description,
	})
	if err != nil {
		return errors.NewDefaultFiberError(err)
	}
	return nil
}
