package bloghandlers

import (
	blogservices "github.com/kuroshibaz/app/blog/services"
	fileservices "github.com/kuroshibaz/app/file/services"
	"github.com/kuroshibaz/lib/validator"
)

type defaultHandler struct {
	cv      validator.CustomValidator
	svc     blogservices.Service
	fileSvc fileservices.Service
}

func NewHandler(cv validator.CustomValidator, svc blogservices.Service, fSvc fileservices.Service) Handler {
	return &defaultHandler{cv: cv, svc: svc, fileSvc: fSvc}
}
