package blogservices

import (
	"github.com/gofiber/fiber/v2"
	blogmodels "github.com/meth-suchatchai/kz-blog-api/app/blog/models"
)

func (svc *defaultService) ListBlog(paginate ...int) (*[]blogmodels.Blog, *fiber.Error) {
	return svc.repo.ListBlog(paginate...)
}
