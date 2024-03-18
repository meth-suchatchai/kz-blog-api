package kzjwt

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kuroshibaz/config"
	"time"
)

type defaultClient struct {
	cfg      *config.JWT
	timezone *time.Location
}

func New(cfg *config.JWT) AuthJWT {
	tz, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Fatal("error load location failed ", err)
	}

	return &defaultClient{
		cfg:      cfg,
		timezone: tz,
	}
}

func (c *defaultClient) generateExpireTime(duration time.Duration) (time.Time, int64) {
	dateTime := time.Now().In(c.timezone).Add(time.Minute * duration)
	return dateTime, dateTime.Unix()
}

func (c *defaultClient) generateTokenClaim(uid int64, exp int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = c.cfg.Issuer
	claims["sub"] = fmt.Sprintf("%d", uid)
	claims["exp"] = exp
	claims["iat"] = time.Now().In(c.timezone).Unix()
	log.Info("claims: ", claims)
	secretKey := []byte(c.cfg.Secret)
	t, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return t, err
}
