package bloghandlers

import (
	"github.com/gofiber/fiber/v2"
	blogmodels "github.com/kuroshibaz/app/blog/models"
	"github.com/kuroshibaz/lib/errors"
)

func (h *defaultHandler) GetBlog(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 0)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	if id == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.ErrBadParameter)
	}

	res, err := h.svc.GetBlog(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	response := &blogmodels.GetBlogResponse{
		Blog: *res,
	}
	return ctx.JSON(response)
}
