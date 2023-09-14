package blogrepositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	blogmodels "github.com/kuroshibaz/app/blog/models"
	"github.com/kuroshibaz/lib/errors"
)

func (repo *defaultRepository) ListBlog(paginate ...int) (*[]blogmodels.Blog, *fiber.Error) {
	result, err := repo.orm.ListBlog(paginate...)
	if err != nil {
		return nil, errors.NewDefaultFiberError(err)
	}

	var blogs []blogmodels.Blog
	for _, blog := range *result {
		tags := []blogmodels.Tag{}

		for _, tag := range blog.Tag {
			t := blogmodels.Tag{
				Id:   int64(tag.ID),
				Name: tag.Name,
				Ord:  tag.Ord,
			}

			tags = append(tags, t)
		}

		m := blogmodels.Blog{
			Id:      int64(blog.ID),
			Name:    blog.Name,
			Content: blog.Content,
			Category: blogmodels.Category{
				Id:   int64(blog.Category.ID),
				Name: blog.Category.Name,
			},
			Tags:     tags,
			ImageURL: blog.Image,
			Slug:     blog.Slug,
			Seo: blogmodels.Seo{
				MetaTitle:       blog.SEO.MetaTitle,
				MetaDescription: blog.SEO.MetaDescription,
			},
			ShortDescription: blog.ShortDescription,
			Views:            blog.Views,
		}
		log.Info("m: ", m)
		blogs = append(blogs, m)
	}

	return &blogs, nil
}
