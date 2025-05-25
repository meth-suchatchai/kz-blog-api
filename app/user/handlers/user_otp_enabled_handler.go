package userhandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	usermodels "github.com/meth-suchatchai/kz-blog-api/app/user/models"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
	coremodels "github.com/meth-suchatchai/kz-blog-api/models"
	"log"
	"strconv"
)

func (h *defaultHandler) OTPEnabled(ctx *fiber.Ctx) error {
	uid := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)
	userId, err := strconv.Atoi(uid)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(errors.ErrInvalidToken)
	}

	var req usermodels.UserOTPEnabledRequest
	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.ErrBadParameter)
	}

	user, vErr := h.userService.UserProfile(int64(userId))
	log.Print("UserProfile: ", user, vErr)
	if vErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(vErr)
	}

	uri, vErr := h.userService.UserOtpToggle(user, req.Enabled)
	log.Print("UserOtpToggle: ", uri, vErr)
	if vErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(vErr)
	}

	return ctx.JSON(coremodels.CreateSuccessResponse(usermodels.UserOTPEnabledResponse{URI: uri}))
}
