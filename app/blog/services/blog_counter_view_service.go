package blogservices

import "github.com/gofiber/fiber/v2"

func (svc *defaultService) CounterView(slug string) *fiber.Error {
	_, err := svc.repo.UpdateViewBySlug(slug)
	if err != nil {
		return err
	}

	return nil
}
