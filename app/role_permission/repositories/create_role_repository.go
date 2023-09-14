package rprepositories

import (
	"github.com/gofiber/fiber/v2"
	rpmodels "github.com/kuroshibaz/app/role_permission/models"
	"github.com/kuroshibaz/lib/errors"
	dbmodels "github.com/kuroshibaz/lib/gormdb/models"
)

func (repo *defaultRepository) CreateRole(data *rpmodels.Role) *fiber.Error {
	err := repo.db.CreateRole(&dbmodels.Role{
		Name:        data.Name,
		Description: data.Description,
	})
	if err != nil {
		return errors.NewDefaultFiberError(err)
	}

	return nil
}
