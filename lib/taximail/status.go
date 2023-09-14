package taximail

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/kuroshibaz/lib/errors"
)

func (c *defaultClient) Status() *fiber.Error {
	res, err := c.Api.R().SetQueryParam("pretty", "true").Get(c.Provide.URL + "/v2/status")
	if err != nil {
		return errors.NewDefaultFiberError(err)
	}
	log.Info(string(res.Body()))

	return nil
}
