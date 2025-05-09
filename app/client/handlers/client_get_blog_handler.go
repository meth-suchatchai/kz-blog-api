package clienthandlers

import (
	"github.com/gofiber/fiber/v2"
	clientmodels "github.com/meth-suchatchai/kz-blog-api/app/client/models"
	constant "github.com/meth-suchatchai/kz-blog-api/const"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
	coremodels "github.com/meth-suchatchai/kz-blog-api/models"
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

	return ctx.JSON(coremodels.CreateSuccessResponse(response))
}
