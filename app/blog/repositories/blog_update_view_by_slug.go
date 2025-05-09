package blogrepositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
)

func (repo *defaultRepository) UpdateViewBySlug(slug string) (int, *fiber.Error) {
	views, err := repo.orm.CountViews(slug)
	if err != nil {
		return 0, errors.NewDefaultFiberError(err)
	}

	return views, nil
}
