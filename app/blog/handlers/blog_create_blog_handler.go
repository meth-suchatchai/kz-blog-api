package bloghandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	blogmodels "github.com/meth-suchatchai/kz-blog-api/app/blog/models"
	constant "github.com/meth-suchatchai/kz-blog-api/const"
	coremodels "github.com/meth-suchatchai/kz-blog-api/models"
)

func (h *defaultHandler) CreateBlog(ctx *fiber.Ctx) error {
	var req blogmodels.CreateBlogRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File[constant.FormFileKey]

	for _, f := range files {
		imagePath, err := h.fileSvc.FileUpload(f, "blog")
		if err != nil {
			log.Warnf("file upload failed: %s", err)
		}
		req.Blog.ImageURL = imagePath
	}

	log.Info("tags: ", req.Tags)
	vErr := h.svc.CreateBlog(&req.Blog)
	if vErr != nil {
		return vErr
	}

	response := &blogmodels.Blog{}
	return ctx.JSON(coremodels.CreateSuccessResponse(response))
}
