package clienthandlers

import (
	"github.com/gofiber/fiber/v2"
	blogmodels "github.com/kuroshibaz/app/blog/models"
	clientmodels "github.com/kuroshibaz/app/client/models"
	constant "github.com/kuroshibaz/const"
	coremodels "github.com/kuroshibaz/models"
)

func (h *defaultHandler) ListBlog(ctx *fiber.Ctx) error {
	page := ctx.QueryInt(constant.QueryPageKey, 0)
	limit := ctx.QueryInt(constant.QueryLimitKey, 0)

	res, err := h.blogService.ListBlog(page, limit)
	if err != nil {
		return err
	}

	blogs := make([]blogmodels.Blog, 0)
	if len(*res) > 0 && res != nil {
		blogs = *res
	}

	response := clientmodels.ListBlogResponse{Blogs: blogs}
	return ctx.JSON(coremodels.CreateSuccessResponse(response))
}
