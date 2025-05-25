package scenehandlers

import "github.com/gofiber/fiber/v2"

type Handler interface {
	CreateScene(ctx *fiber.Ctx) error
	UpdateStatusScene(ctx *fiber.Ctx) error
}
