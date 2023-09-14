package blogservices

import (
	"github.com/gofiber/fiber/v2"
	blogmodels "github.com/kuroshibaz/app/blog/models"
)

func (svc *defaultService) CreateBlog(data *blogmodels.Blog) *fiber.Error {
	return svc.repo.CreateBlog(data)
}
