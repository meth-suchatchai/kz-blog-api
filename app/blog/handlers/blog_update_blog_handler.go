package bloghandlers

import (
	"github.com/gofiber/fiber/v2"
	coremodels "github.com/meth-suchatchai/kz-blog-api/models"
)

func (h *defaultHandler) UpdateBlog(ctx *fiber.Ctx) error {
	return ctx.JSON(coremodels.CreateSuccessResponse(""))
}
