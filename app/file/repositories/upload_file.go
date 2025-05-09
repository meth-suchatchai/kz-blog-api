package filerepositories

import (
	"github.com/gofiber/fiber/v2"
	filemodels "github.com/meth-suchatchai/kz-blog-api/app/file/models"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
	"github.com/meth-suchatchai/kz-blog-api/lib/kzobjectstorage"
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
