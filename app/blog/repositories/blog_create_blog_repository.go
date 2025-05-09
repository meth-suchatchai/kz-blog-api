package blogrepositories

import (
	"github.com/gofiber/fiber/v2"
	blogmodels "github.com/meth-suchatchai/kz-blog-api/app/blog/models"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
	dbmodels "github.com/meth-suchatchai/kz-blog-api/lib/gormdb/models"
)

func (repo *defaultRepository) CreateBlog(data *blogmodels.Blog) *fiber.Error {
	tags := []dbmodels.Tag{}
	for _, b := range data.Tags {
		m := dbmodels.Tag{
			Name: b.Name,
			Ord:  b.Ord,
		}
		tags = append(tags, m)
	}

	err := repo.orm.CreateBlog(&dbmodels.Blog{
		Name:             data.Name,
		Content:          data.Content,
		ShortDescription: data.ShortDescription,
		Image:            data.ImageURL,
		Tag:              tags,
		CategoryId:       int(data.Category.Id),
		SEO: dbmodels.SEO{
			MetaTitle:       data.Seo.MetaTitle,
			MetaDescription: data.Seo.MetaDescription,
		},
	})
	if err != nil {
		return errors.NewDefaultFiberError(err)
	}

	return nil
}
