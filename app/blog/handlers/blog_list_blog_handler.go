package bloghandlers

import (
	"github.com/gofiber/fiber/v2"
	blogmodels "github.com/meth-suchatchai/kz-blog-api/app/blog/models"
	constant "github.com/meth-suchatchai/kz-blog-api/const"
	coremodels "github.com/meth-suchatchai/kz-blog-api/models"
)

func (h *defaultHandler) ListBlog(ctx *fiber.Ctx) error {
	page := ctx.QueryInt(constant.QueryPageKey, 0)
	limit := ctx.QueryInt(constant.QueryLimitKey, 0)

	blogs, err := h.svc.ListBlog(page, limit)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	newBlogs := make([]blogmodels.Blog, 0)
	if blogs != nil && len(*blogs) > 0 {
		newBlogs = *blogs
	}
	response := &blogmodels.ListBlogResponse{Blogs: newBlogs}

	//return ctx.Status(fiber.StatusInternalServerError).JSON(errors.NewDefaultFiberMessageError("test"))

	return ctx.JSON(coremodels.CreateSuccessResponse(response))
}
