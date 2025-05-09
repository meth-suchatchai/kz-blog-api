package rpservices

import (
	"github.com/gofiber/fiber/v2"
	rpmodels "github.com/meth-suchatchai/kz-blog-api/app/role_permission/models"
)

func (svc *defaultService) CreateRole(data *rpmodels.Role) *fiber.Error {
	return svc.rp.CreateRole(data)
}
