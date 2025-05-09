package filerepositories

import (
	"github.com/gofiber/fiber/v2"
	filemodels "github.com/meth-suchatchai/kz-blog-api/app/file/models"
)

type Repository interface {
	UploadFile(file *filemodels.UploadFileData) (string, *fiber.Error)
}
