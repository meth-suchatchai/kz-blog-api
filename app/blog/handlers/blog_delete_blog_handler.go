package bloghandlers

import (
	"github.com/gofiber/fiber/v2"
	blogmodels "github.com/meth-suchatchai/kz-blog-api/app/blog/models"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
	coremodels "github.com/meth-suchatchai/kz-blog-api/models"
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

	response := &blogmodels.DeleteBlogResponse{}
	return ctx.JSON(coremodels.CreateSuccessResponse(response))
}
