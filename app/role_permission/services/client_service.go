package rpservices

import (
	"github.com/gofiber/fiber/v2"
	rpmodels "github.com/meth-suchatchai/kz-blog-api/app/role_permission/models"
)

type Service interface {
	CreateRole(data *rpmodels.Role) *fiber.Error
	CreatePermission(data *rpmodels.Permission) *fiber.Error
	RolePermission() (*[]rpmodels.RolePermission, *fiber.Error)
}
