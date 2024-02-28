package userhandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	usermodels "github.com/kuroshibaz/app/user/models"
	"github.com/kuroshibaz/lib/errors"
	coremodels "github.com/kuroshibaz/models"
	"strconv"
)

func (h *defaultHandler) Profile(ctx *fiber.Ctx) error {
	uid := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)
	userId, err := strconv.Atoi(uid)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(errors.ErrInvalidToken)
	}

	user, err := h.userService.UserProfile(int64(userId))

	response := usermodels.ProfileResponse{UserProfile: usermodels.UserProfile{
		Id:               user.Id,
		Name:             user.Name,
		MobileNumber:     user.MobileNumber,
		TwoFactorEnabled: user.IsTFA,
	}}

	return ctx.JSON(coremodels.CreateSuccessResponse(response))
}
