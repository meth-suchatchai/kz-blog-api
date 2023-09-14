package kzobjectstorage

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
)

type Options struct {
	Endpoint        string
	AccessKeyId     string
	SecretAccessKey string
	UseSSL          bool
	Region          string
}

type defaultClient struct {
	opts        *Options
	minioClient *minio.Client
	bucketName  string
}

func NewClient(opts *Options) (StorageMinio, error) {
	if opts.AccessKeyId == "" || opts.SecretAccessKey == "" {
		return nil, errors.New("access key or secret key not found")
	}

	if opts.Region == "" {
		log.Warn("region not set will use default us-east-1")
		opts.Region = "us-east-1"
	}

	minioClient, err := minio.New(opts.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(opts.AccessKeyId, opts.SecretAccessKey, ""),
		Secure: opts.UseSSL,
	})

	if err != nil {
		return nil, err
	}

	return &defaultClient{opts: opts, minioClient: minioClient}, nil
}

func (c *defaultClient) Minio() *minio.Client {
	return c.minioClient
}
