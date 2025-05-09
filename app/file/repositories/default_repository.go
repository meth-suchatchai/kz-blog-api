package filerepositories

import "github.com/meth-suchatchai/kz-blog-api/lib/kzobjectstorage"

type defaultRepository struct {
	storageBucket kzobjectstorage.StorageBucket
}

func NewRepository(storageBucket kzobjectstorage.StorageBucket) Repository {
	return &defaultRepository{storageBucket: storageBucket}
}
