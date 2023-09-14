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
	accessTokenTime := c.generateExpireTime(time.Minute * time.Duration(c.cfg.Expire))
	refreshTokenTime := c.generateExpireTime(time.Minute * time.Duration(c.cfg.RefreshExpire))

	accessToken, err := c.generateTokenClaim(data.Id, accessTokenTime)
	if err != nil {
		return nil, fiber.NewError(errors.ErrCodeInvalidJWT, err.Error())
	}
	refreshToken, err := c.generateTokenClaim(data.Id, refreshTokenTime)
	if err != nil {
		return nil, fiber.NewError(errors.ErrCodeInvalidJWT, err.Error())
	}

	return &AccessToken{
		AccessToken:        accessToken,
		AccessTokenExpire:  accessTokenTime,
		RefreshToken:       refreshToken,
		RefreshTokenExpire: refreshTokenTime,
		Domain:             c.cfg.Domain,
	}, nil
}
