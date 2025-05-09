package blogservices

import (
	"github.com/gofiber/fiber/v2"
	blogmodels "github.com/meth-suchatchai/kz-blog-api/app/blog/models"
)

func (svc *defaultService) GetBlog(id int) (*blogmodels.Blog, *fiber.Error) {
	return svc.repo.GetBlogById(id)
}
