package scenehandlers

import (
	sceneservices "github.com/meth-suchatchai/kz-blog-api/app/scene/services"
	"github.com/meth-suchatchai/kz-blog-api/lib/validator"
)

type defaultHandler struct {
	cv  validator.CustomValidator
	svc sceneservices.Service
}

func NewHandler(cv validator.CustomValidator, svc sceneservices.Service) Handler {
	return &defaultHandler{cv: cv, svc: svc}
}
