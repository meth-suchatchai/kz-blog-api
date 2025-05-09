package rprepositories

import (
	"github.com/gofiber/fiber/v2"
	rpmodels "github.com/meth-suchatchai/kz-blog-api/app/role_permission/models"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
	dbmodels "github.com/meth-suchatchai/kz-blog-api/lib/gormdb/models"
)

func (repo *defaultRepository) CreatePermission(data *rpmodels.Permission) *fiber.Error {
	err := repo.orm.CreatePermission(&dbmodels.Permission{
		Name:        data.Name,
		Code:        data.Code,
		Description: data.Description,
	})
	if err != nil {
		return errors.NewDefaultFiberError(err)
	}
	return nil
}
