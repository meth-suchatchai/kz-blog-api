package etcdservices

import (
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	GetDataByKey(key string, output interface{}) *fiber.Error
	GetDataStringByKey(key string) (string, *fiber.Error)
}
