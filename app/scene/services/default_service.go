package sceneservices

import (
	scenerepositories "github.com/meth-suchatchai/kz-blog-api/app/scene/repositories"
)

type defaultService struct {
	rp scenerepositories.Repository
}

func NewService(rp scenerepositories.Repository) Service {
	return &defaultService{rp: rp}
}
