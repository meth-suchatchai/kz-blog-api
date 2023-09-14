package blogservices

import (
	"github.com/gofiber/fiber/v2"
	blogmodels "github.com/kuroshibaz/app/blog/models"
)

func (svc *defaultService) GetBlog(id int) (*blogmodels.Blog, *fiber.Error) {
	return svc.repo.GetBlogById(id)
}
