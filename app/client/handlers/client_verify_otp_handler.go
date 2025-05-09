package clienthandlers

import (
	"github.com/gofiber/fiber/v2"
	clientmodels "github.com/meth-suchatchai/kz-blog-api/app/client/models"
	coremodels "github.com/meth-suchatchai/kz-blog-api/models"
)

func (h *defaultHandler) VerifyOTP(ctx *fiber.Ctx) error {
	var req clientmodels.VerifyOTPRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := h.clientService.VerifyOTP(req.VerifyOTPData)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	response := clientmodels.VerifyOTPResponse{}
	return ctx.JSON(coremodels.CreateSuccessResponse(response))
}
