package scenehandlers

import (
	"github.com/gofiber/fiber/v2"
	scenemodels "github.com/meth-suchatchai/kz-blog-api/app/scene/models"
	"strconv"
)

func (h *defaultHandler) UpdateStatusScene(ctx *fiber.Ctx) error {
	id := ctx.Get("id")

	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	var req scenemodels.UpdateSceneStatusRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	vErr := h.svc.UpdateSceneStatus(uint(uid), req.Status)
	if vErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ctx.JSON(scenemodels.UpdateSceneResponse{})
}
