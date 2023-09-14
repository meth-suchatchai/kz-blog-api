package bloghandlers

import "github.com/gofiber/fiber/v2"

type Handler interface {
	ListBlog(ctx *fiber.Ctx) error
	GetBlog(ctx *fiber.Ctx) error
	CreateBlog(ctx *fiber.Ctx) error
	UpdateBlog(ctx *fiber.Ctx) error
	DeleteBlog(ctx *fiber.Ctx) error
}
