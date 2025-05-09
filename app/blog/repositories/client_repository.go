package blogrepositories

import (
	"github.com/gofiber/fiber/v2"
	blogmodels "github.com/meth-suchatchai/kz-blog-api/app/blog/models"
)

type Repository interface {
	CreateBlog(data *blogmodels.Blog) *fiber.Error
	GetBlogById(id int) (*blogmodels.Blog, *fiber.Error)
	GetBlogBySlug(slug string) (*blogmodels.Blog, *fiber.Error)
	DeleteBlog(id int) *fiber.Error
	ListBlog(paginate ...int) (*[]blogmodels.Blog, *fiber.Error)
	UpdateViewBySlug(slug string) (int, *fiber.Error)
}
