package userrepositories

import "github.com/kuroshibaz/lib/gormdb"

type defaultRepository struct {
	cli *gormdb.DB
}

func NewRepository(cli *gormdb.DB) Repository {
	return &defaultRepository{
		cli,
	}
}
