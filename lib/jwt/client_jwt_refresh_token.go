package kzjwt

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kuroshibaz/lib/errors"
	"strconv"
	"time"
)

func (c *defaultClient) JwtRefreshToken(refreshToken string) (*AccessToken, *fiber.Error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(c.cfg.Secret), nil
	})

	if err != nil {
		return nil, fiber.NewError(errors.ErrCodeInvalidJWT, err.Error())
	}
	if !token.Valid {
		return nil, fiber.NewError(errors.ErrCodeInvalidJWT, "invalid or expired refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fiber.NewError(errors.ErrCodeInvalidJWT, "failed to parse token")
	}

	u, ok := claims["sub"].(string)
	if !ok {
		return nil, fiber.NewError(errors.ErrCodeInvalidJWT, "failed to extract uid")
	}

	uid, err := strconv.Atoi(u)
	if err != nil {
		return nil, fiber.NewError(errors.ErrCodeInvalidJWT, err.Error())
	}

	extentTime := c.generateExpireTime(time.Minute * time.Duration(c.cfg.Expire))
	newRefreshTime := c.generateExpireTime(time.Minute * time.Duration(c.cfg.RefreshExpire))

	accessToken, vErr := c.generateTokenClaim(int64(uid), extentTime)
	if vErr != nil {
		return nil, fiber.NewError(errors.ErrCodeInvalidJWT, vErr.Error())
	}

	newRefreshToken, vErr := c.generateTokenClaim(int64(uid), newRefreshTime)
	if vErr != nil {
		return nil, fiber.NewError(errors.ErrCodeInvalidJWT, vErr.Error())
	}

	return &AccessToken{
		AccessToken:        accessToken,
		AccessTokenExpire:  extentTime,
		RefreshToken:       newRefreshToken,
		RefreshTokenExpire: newRefreshTime,
		Domain:             c.cfg.Domain,
	}, nil
}
