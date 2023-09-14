package blogrepositories

import "github.com/kuroshibaz/lib/gormdb"

type defaultRepository struct {
	orm *gormdb.DB
}

func NewRepository(db *gormdb.DB) Repository {
	return &defaultRepository{orm: db}
}
