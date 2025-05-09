package rphandlers

import (
	"github.com/gofiber/fiber/v2"
	rpmodels "github.com/meth-suchatchai/kz-blog-api/app/role_permission/models"
)

func (h *defaultHandler) CreatePermission(ctx *fiber.Ctx) error {
	var req rpmodels.CreatePermissionRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := h.svc.CreatePermission(&req.Permission)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ctx.JSON(rpmodels.CreatePermissionResponse{})
}
