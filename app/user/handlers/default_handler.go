package userhandlers

import (
	userservices "github.com/meth-suchatchai/kz-blog-api/app/user/services"
	"github.com/meth-suchatchai/kz-blog-api/lib/validator"
)

type defaultHandler struct {
	validate    validator.CustomValidator
	userService userservices.Service
}

func NewHandler(validate validator.CustomValidator, userService userservices.Service) Handler {
	return &defaultHandler{
		validate:    validate,
		userService: userService,
	}
}
