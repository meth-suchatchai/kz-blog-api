package fileservices

import filerepositories "github.com/meth-suchatchai/kz-blog-api/app/file/repositories"

type defaultService struct {
	fileRepo filerepositories.Repository
}

func NewService(fileRepo filerepositories.Repository) Service {
	return &defaultService{fileRepo: fileRepo}
}
