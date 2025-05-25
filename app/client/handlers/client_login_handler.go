package clienthandlers

import (
	"github.com/gofiber/fiber/v2"
	clientmodels "github.com/meth-suchatchai/kz-blog-api/app/client/models"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
	coremodels "github.com/meth-suchatchai/kz-blog-api/models"
)

func (h *defaultHandler) Login(ctx *fiber.Ctx) error {
	var req clientmodels.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.NewDefaultFiberError(err))
	}

	user, auth, err := h.clientService.Login(req.LoginData)
	if err != nil {
		if err.Code == 203 {
			return ctx.Status(fiber.StatusOK).JSON(coremodels.SuccessResponse{
				Code:    err.Code,
				Message: "2FA required",
				Data:    nil,
			})
		}
		return err
	}

	response := clientmodels.LoginResponse{
		Authentication: clientmodels.Authentication{
			AccessToken:        auth.AccessToken,
			AccessTokenExpire:  auth.AccessTokenExpire,
			RefreshToken:       auth.RefreshToken,
			RefreshTokenExpire: auth.RefreshTokenExpire,
			Domain:             auth.Domain,
		},
		LoginUser: clientmodels.LoginUser{
			Id:           user.Id,
			Name:         user.Name,
			MobileNumber: user.MobileNumber,
		},
	}

	return ctx.JSON(coremodels.CreateSuccessResponse(response))
}
