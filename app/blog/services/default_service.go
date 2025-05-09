package blogservices

import blogrepositories "github.com/meth-suchatchai/kz-blog-api/app/blog/repositories"

type defaultService struct {
	repo blogrepositories.Repository
}

func NewService(repo blogrepositories.Repository) Service {
	return &defaultService{repo: repo}
}
