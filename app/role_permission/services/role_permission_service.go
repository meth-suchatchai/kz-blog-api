package rpservices

import (
	"github.com/gofiber/fiber/v2"
	rpmodels "github.com/kuroshibaz/app/role_permission/models"
)

func (svc *defaultService) RolePermission() (*[]rpmodels.RolePermission, *fiber.Error) {
	return svc.rp.RolePermission()
}
