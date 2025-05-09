package blogservices

import (
	"github.com/gofiber/fiber/v2"
	blogmodels "github.com/meth-suchatchai/kz-blog-api/app/blog/models"
)

type Service interface {
	CreateBlog(data *blogmodels.Blog) *fiber.Error
	GetBlog(id int) (*blogmodels.Blog, *fiber.Error)
	GetBlogBySlug(slug string) (*blogmodels.Blog, *fiber.Error)
	DeleteBlog(id int) *fiber.Error
	ListBlog(paginate ...int) (*[]blogmodels.Blog, *fiber.Error)
	CounterView(slug string) *fiber.Error
}
