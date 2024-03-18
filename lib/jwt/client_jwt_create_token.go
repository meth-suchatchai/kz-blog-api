package kzjwt

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	usermodels "github.com/kuroshibaz/app/user/models"
	"github.com/kuroshibaz/lib/errors"
	"time"
)

func (c *defaultClient) JwtCreateToken(data *usermodels.User) (*AccessToken, *fiber.Error) {
	log.Info("config: ", c.cfg)
	accessTokenExpireDate, accessTokenExpire := c.generateExpireTime(time.Minute * time.Duration(c.cfg.Expire))
	refreshTokenExpireDate, refreshTokenExpire := c.generateExpireTime(time.Minute * time.Duration(c.cfg.RefreshExpire))

	accessToken, err := c.generateTokenClaim(data.Id, accessTokenExpire)
	if err != nil {
		return nil, fiber.NewError(errors.ErrCodeInvalidJWT, err.Error())
	}
	refreshToken, err := c.generateTokenClaim(data.Id, refreshTokenExpire)
	if err != nil {
		return nil, fiber.NewError(errors.ErrCodeInvalidJWT, err.Error())
	}

	return &AccessToken{
		AccessToken:        accessToken,
		AccessTokenExpire:  accessTokenExpireDate,
		RefreshToken:       refreshToken,
		RefreshTokenExpire: refreshTokenExpireDate,
		Domain:             c.cfg.Domain,
	}, nil
}
