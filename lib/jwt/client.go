package kzjwt

import (
	"github.com/gofiber/fiber/v2"
	usermodels "github.com/meth-suchatchai/kz-blog-api/app/user/models"
)

type AuthJWT interface {
	JwtCreateToken(data *usermodels.User) (*AccessToken, *fiber.Error)
	JwtRefreshToken(refreshToken string) (*AccessToken, *fiber.Error)
}
