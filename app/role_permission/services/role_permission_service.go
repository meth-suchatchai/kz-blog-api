package rpservices

import (
	"github.com/gofiber/fiber/v2"
	rpmodels "github.com/meth-suchatchai/kz-blog-api/app/role_permission/models"
)

func (svc *defaultService) RolePermission() (*[]rpmodels.RolePermission, *fiber.Error) {
	return svc.rp.RolePermission()
}
