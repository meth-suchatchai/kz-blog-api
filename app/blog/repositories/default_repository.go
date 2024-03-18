package blogrepositories

import "github.com/kuroshibaz/lib/gormdb"

type defaultRepository struct {
	orm gormdb.Client
}

func NewRepository(orm gormdb.Client) Repository {
	return &defaultRepository{orm: orm}
}
