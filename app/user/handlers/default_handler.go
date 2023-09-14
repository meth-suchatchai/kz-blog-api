package userhandlers

import (
	userservices "github.com/kuroshibaz/app/user/services"
	"github.com/kuroshibaz/lib/validator"
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
