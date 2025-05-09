package rphandlers

import (
	"github.com/gofiber/fiber/v2"
	rpmodels "github.com/meth-suchatchai/kz-blog-api/app/role_permission/models"
)

func (h *defaultHandler) CreateRole(ctx *fiber.Ctx) error {
	var req rpmodels.CreateRoleRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := h.svc.CreateRole(&req.Role)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ctx.JSON(rpmodels.CreateRoleResponse{})
}
