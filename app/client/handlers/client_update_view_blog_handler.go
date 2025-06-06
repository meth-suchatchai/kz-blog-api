package clienthandlers

import (
	"github.com/gofiber/fiber/v2"
	constant "github.com/meth-suchatchai/kz-blog-api/const"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
	coremodels "github.com/meth-suchatchai/kz-blog-api/models"
)

func (h *defaultHandler) UpdateViewBlog(ctx *fiber.Ctx) error {
	slug := ctx.Params(constant.QuerySlugKey, "")
	if slug == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.ErrBadParameter)
	}

	err := h.blogService.CounterView(slug)
	if err != nil {
		return err
	}

	return ctx.JSON(coremodels.CreateSuccessResponse(""))
}
