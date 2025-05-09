package userrepositories

import "github.com/meth-suchatchai/kz-blog-api/lib/gormdb"

type defaultRepository struct {
	orm gormdb.Client
}

func NewRepository(orm gormdb.Client) Repository {
	return &defaultRepository{
		orm: orm,
	}
}
