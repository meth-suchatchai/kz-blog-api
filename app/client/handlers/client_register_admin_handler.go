package clienthandlers

import (
	"github.com/gofiber/fiber/v2"
	clientmodels "github.com/kuroshibaz/app/client/models"
	coremodels "github.com/kuroshibaz/models"
)

func (h *defaultHandler) RegisterAdmin(ctx *fiber.Ctx) error {
	registerValue := ctx.Query("allow_register", "")
	if registerValue == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400, "no parameter found"))
	}

	constantAllowRegisterValue, err := h.etcdService.GetDataStringByKey("allow_register")
	if err != nil {
		return err
	}

	if constantAllowRegisterValue != registerValue {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.NewError(403, "pass code incorrect. Register not allow"))
	}

	var req clientmodels.RegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	userVerify, vErr := h.clientService.Register(req.RegisterData)
	if vErr != nil {
		return vErr
	}

	response := clientmodels.RegisterResponse{
		OTPReferenceNumber: userVerify.OTPReferenceNumber,
	}

	return ctx.JSON(coremodels.CreateSuccessResponse(response))
}
