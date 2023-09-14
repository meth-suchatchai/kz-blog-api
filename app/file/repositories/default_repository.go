package filerepositories

import "github.com/kuroshibaz/lib/kzobjectstorage"

type defaultRepository struct {
	storageBucket kzobjectstorage.StorageBucket
}

func NewRepository(storageBucket kzobjectstorage.StorageBucket) Repository {
	return &defaultRepository{storageBucket: storageBucket}
}
