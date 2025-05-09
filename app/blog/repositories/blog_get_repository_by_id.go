package blogrepositories

import (
	"github.com/gofiber/fiber/v2"
	blogmodels "github.com/meth-suchatchai/kz-blog-api/app/blog/models"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
)

func (repo *defaultRepository) GetBlogById(id int) (*blogmodels.Blog, *fiber.Error) {
	blog, err := repo.orm.GetBlogById(uint(id))
	if err != nil {
		return nil, errors.NewDefaultFiberError(err)
	}

	tags := make([]blogmodels.Tag, 0)
	for _, b := range blog.Tag {
		m := blogmodels.Tag{
			Id:   int64(b.ID),
			Name: b.Name,
			Ord:  b.Ord,
		}
		tags = append(tags, m)
	}

	response := &blogmodels.Blog{
		Id:       int64(blog.ID),
		Name:     blog.Name,
		Content:  blog.Content,
		ImageURL: blog.Image,
		Slug:     blog.Slug,
		Category: blogmodels.Category{
			Id:   int64(blog.Category.ID),
			Name: blog.Category.Name,
		},
		Tags: tags,
		Seo: blogmodels.Seo{
			MetaTitle:       blog.SEO.MetaTitle,
			MetaDescription: blog.SEO.MetaDescription,
		},
	}
	return response, nil
}
