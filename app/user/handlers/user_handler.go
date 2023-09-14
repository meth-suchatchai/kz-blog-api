package userhandlers

import "github.com/gofiber/fiber/v2"

type Handler interface {
	ListUser(ctx *fiber.Ctx) error
	Profile(ctx *fiber.Ctx) error
}
