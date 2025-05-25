package scenehandlers

import (
	"github.com/gofiber/fiber/v2"
	scenemodels "github.com/meth-suchatchai/kz-blog-api/app/scene/models"
)

func (h *defaultHandler) CreateScene(ctx *fiber.Ctx) error {
	var req scenemodels.CreateSceneRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	res, err := h.svc.CreateScene(req.CreateSceneData)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ctx.JSON(scenemodels.CreateSceneResponse{Scene: *res})
}
