package fileservices

import (
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
)

type Service interface {
	FileUpload(multipart *multipart.FileHeader, dir ...string) (string, *fiber.Error)
}
