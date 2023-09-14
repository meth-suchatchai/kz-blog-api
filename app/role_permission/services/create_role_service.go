package rpservices

import (
	"github.com/gofiber/fiber/v2"
	rpmodels "github.com/kuroshibaz/app/role_permission/models"
)

func (svc *defaultService) CreateRole(data *rpmodels.Role) *fiber.Error {
	return svc.rp.CreateRole(data)
}
