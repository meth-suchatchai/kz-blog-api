package clienthandlers

import (
	"github.com/gofiber/fiber/v2"
	clientmodels "github.com/kuroshibaz/app/client/models"
	constant "github.com/kuroshibaz/const"
	"github.com/kuroshibaz/lib/errors"
)

func (h *defaultHandler) GetBlog(ctx *fiber.Ctx) error {
	slug := ctx.Params(constant.QuerySlugKey, "")
	if slug == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(errors.ErrBadParameter)
	}

	blog, err := h.blogService.GetBlogBySlug(slug)
	if err != nil {
		return err
	}

	response := clientmodels.GetBlogResponse{Blog: *blog}

	return ctx.JSON(clientmodels.CreateSuccessResponse(response))
}
