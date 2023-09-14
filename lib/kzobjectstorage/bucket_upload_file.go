package kzobjectstorage

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"github.com/minio/minio-go/v7"
	"os"
)

type UploadPutObjectOption struct {
	Filename    string
	Directory   string
	ContentType string
	Size        int64
	File        *os.File
}

func (c *defaultBucket) UploadFile(object *UploadPutObjectOption) (string, error) {
	pathFile := c.getDirectoryPath(object.Directory, object.Filename)
	log.Info("file", object.File.Name())
	_, err := c.minioClient.FPutObject(
		context.TODO(),
		c.bucketName,
		pathFile,
		object.File.Name(),
		minio.PutObjectOptions{
			ContentType: object.ContentType,
		})
	if err != nil {
		return "", err
	}

	return c.getPublicPath(pathFile), nil
}
