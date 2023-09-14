package rphandlers

import "github.com/gofiber/fiber/v2"

type Handler interface {
	ListRolePermission(ctx *fiber.Ctx) error
	CreatePermission(ctx *fiber.Ctx) error
	CreateRole(ctx *fiber.Ctx) error
}
