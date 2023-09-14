package clienthandlers

import "github.com/gofiber/fiber/v2"

type Handler interface {
	Login(ctx *fiber.Ctx) error
	RegisterAdmin(ctx *fiber.Ctx) error
	VerifyOTP(ctx *fiber.Ctx) error
	ListBlog(ctx *fiber.Ctx) error
	GetBlog(ctx *fiber.Ctx) error
	UpdateViewBlog(ctx *fiber.Ctx) error
}
