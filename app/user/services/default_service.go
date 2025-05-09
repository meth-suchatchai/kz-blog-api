package userservices

import (
	userrepositories "github.com/meth-suchatchai/kz-blog-api/app/user/repositories"
)

type defaultService struct {
	userRepo userrepositories.Repository
}

func NewService(userRepo userrepositories.Repository) Service {
	return &defaultService{
		userRepo: userRepo,
	}
}
