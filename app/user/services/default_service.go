package userservices

import (
	userrepositories "github.com/meth-suchatchai/kz-blog-api/app/user/repositories"
	"github.com/meth-suchatchai/kz-blog-api/lib/totp"
)

type defaultService struct {
	userRepo userrepositories.Repository
	totp     totp.Client
}

func NewService(userRepo userrepositories.Repository, totp totp.Client) Service {
	return &defaultService{
		userRepo: userRepo,
		totp:     totp,
	}
}
