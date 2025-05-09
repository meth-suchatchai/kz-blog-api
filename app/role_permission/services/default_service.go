package rpservices

import rprepositories "github.com/meth-suchatchai/kz-blog-api/app/role_permission/repositories"

type defaultService struct {
	rp rprepositories.Repository
}

func NewService(rp rprepositories.Repository) Service {
	return &defaultService{rp: rp}
}
