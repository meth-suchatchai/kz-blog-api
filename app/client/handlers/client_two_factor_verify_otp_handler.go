package clienthandlers

import (
	"github.com/gofiber/fiber/v2"
	clientmodels "github.com/kuroshibaz/app/client/models"
)

func (h *defaultHandler) TwoFactorVerifyOtp(ctx *fiber.Ctx) error {
	var req clientmodels.TwoFactorVerifyRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := h.clientService.TwoFactorVerify(req.TwoFactorVerifyData)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return nil
}
