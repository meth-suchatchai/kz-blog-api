package kzobjectstorage

import "github.com/minio/minio-go/v7"

type StorageMinio interface {
	Minio() *minio.Client
}
