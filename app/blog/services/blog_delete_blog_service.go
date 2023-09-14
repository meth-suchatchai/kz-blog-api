package blogservices

import "github.com/gofiber/fiber/v2"

func (svc *defaultService) DeleteBlog(id int) *fiber.Error {
	return svc.repo.DeleteBlog(id)
}
