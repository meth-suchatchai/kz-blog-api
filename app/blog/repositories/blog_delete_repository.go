package blogrepositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kuroshibaz/lib/errors"
)

func (repo *defaultRepository) DeleteBlog(id int) *fiber.Error {
	err := repo.orm.DeleteBlog(uint(id))
	if err != nil {
		return errors.NewDefaultFiberError(err)
	}

	return nil
}
