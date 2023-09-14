package rphandlers

import (
	"github.com/gofiber/fiber/v2"
	rpmodels "github.com/kuroshibaz/app/role_permission/models"
)

func (h *defaultHandler) ListRolePermission(ctx *fiber.Ctx) error {
	//var req rpmodels.RolePermissionRequest
	//if err := ctx.BodyParser(&req); err != nil {
	//	return ctx.Status(fiber.StatusBadRequest).JSON(err)
	//}

	result, err := h.svc.RolePermission()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	rp := make([]rpmodels.RolePermission, 0)
	if result != nil && len(*result) > 0 {
		rp = *result
	}

	response := rpmodels.RolePermissionResponse{RolePermissions: rp}
	return ctx.JSON(response)
}
