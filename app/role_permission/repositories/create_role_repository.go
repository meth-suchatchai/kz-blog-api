package rprepositories

import (
	"github.com/gofiber/fiber/v2"
	rpmodels "github.com/meth-suchatchai/kz-blog-api/app/role_permission/models"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
	dbmodels "github.com/meth-suchatchai/kz-blog-api/lib/gormdb/models"
)

func (repo *defaultRepository) CreateRole(data *rpmodels.Role) *fiber.Error {
	err := repo.orm.CreateRole(&dbmodels.Role{
		Name:        data.Name,
		Description: data.Description,
	})
	if err != nil {
		return errors.NewDefaultFiberError(err)
	}

	return nil
}
