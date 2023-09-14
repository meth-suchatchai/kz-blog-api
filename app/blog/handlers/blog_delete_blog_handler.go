package bloghandlers

import (
	"github.com/gofiber/fiber/v2"
	blogmodels "github.com/kuroshibaz/app/blog/models"
	"github.com/kuroshibaz/lib/errors"
)

func (h *defaultHandler) DeleteBlog(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 0)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	if id == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.ErrBadParameter)
	}

	err = h.svc.DeleteBlog(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ctx.JSON(&blogmodels.DeleteBlogResponse{})
}
