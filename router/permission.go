package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kuroshibaz/lib/errors"
	"github.com/kuroshibaz/lib/gormdb"
	"strconv"
)

type Permission interface {
	CheckPermission(permission string) fiber.Handler
}

type defaultPermission struct {
	orm *gormdb.DB
}

func NewPermission(db *gormdb.DB) Permission {
	return &defaultPermission{orm: db}
}

func (pm *defaultPermission) CheckPermission(permission string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		uid := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)
		userId, err := strconv.Atoi(uid)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(errors.ErrInvalidToken)
		}

		p, err := pm.orm.GetPermission(permission)
		if err != nil {
			return ctx.Status(fiber.StatusForbidden).JSON(errors.ErrPermissionDenied)
		}

		hasPermission := pm.orm.GetUserPermission(uint(userId), p.ID)
		if !hasPermission {
			return ctx.Status(fiber.StatusForbidden).JSON(errors.ErrPermissionDenied)
		}

		return ctx.Next()
	}
}
