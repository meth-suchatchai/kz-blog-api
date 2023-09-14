package filerepositories

import (
	"github.com/gofiber/fiber/v2"
	filemodels "github.com/kuroshibaz/app/file/models"
	"github.com/kuroshibaz/lib/errors"
	"github.com/kuroshibaz/lib/kzobjectstorage"
	"log"
)

func (c *defaultRepository) UploadFile(data *filemodels.UploadFileData) (string, *fiber.Error) {
	path, vErr := c.storageBucket.UploadFile(&kzobjectstorage.UploadPutObjectOption{
		Filename:    data.Filename,
		Directory:   data.Directory,
		Size:        data.Size,
		ContentType: data.MimeType,
		File:        data.File,
	})
	if vErr != nil {
		return "", errors.NewDefaultFiberError(vErr)
	}

	log.Print("upload success : ", path)
	return path, nil
}
