package rprepositories

import "github.com/kuroshibaz/lib/gormdb"

type defaultRepository struct {
	db *gormdb.DB
}

func NewRepository(db *gormdb.DB) Repository {
	return &defaultRepository{db: db}
}
