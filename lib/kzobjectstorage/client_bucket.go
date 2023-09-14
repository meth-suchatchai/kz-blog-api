package kzobjectstorage

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
)

type defaultBucket struct {
	bucketName  string
	endpoint    string
	minioClient *minio.Client
}

type StorageBucket interface {
	UploadFile(object *UploadPutObjectOption) (string, error)
	ListBucket() ([]string, error)
}

func NewSelectBucket(bucketName, endpoint, region string, minioClient *minio.Client) (StorageBucket, error) {
	if bucketName == "" {
		return nil, errors.New("bucket not found")
	}
	if endpoint == "" {
		return nil, errors.New("endpoint not found")
	}
	dfb := defaultBucket{
		bucketName:  bucketName,
		endpoint:    endpoint,
		minioClient: minioClient,
	}
	if exist, err := dfb.minioClient.BucketExists(context.TODO(), bucketName); err != nil {
		return nil, err
	} else {
		if !exist {
			log.Warn("bucket not exist create a new one")
			if err = dfb.minioClient.MakeBucket(context.TODO(), bucketName, minio.MakeBucketOptions{
				Region: region,
			}); err != nil {
				return nil, err
			}
		}
	}

	return &dfb, nil
}

func (c *defaultBucket) getDirectoryPath(directory, fileName string) string {
	if directory == "" {
		return fileName
	}

	return fmt.Sprintf("%v/%v", directory, fileName)
}

func (c *defaultBucket) getPublicPath(filePath string) string {
	return fmt.Sprintf("%s/%s/%s", c.endpoint, c.bucketName, filePath)
}
