package bloghandlers

import (
	blogservices "github.com/meth-suchatchai/kz-blog-api/app/blog/services"
	fileservices "github.com/meth-suchatchai/kz-blog-api/app/file/services"
	"github.com/meth-suchatchai/kz-blog-api/lib/validator"
)

type defaultHandler struct {
	cv      validator.CustomValidator
	svc     blogservices.Service
	fileSvc fileservices.Service
}

func NewHandler(cv validator.CustomValidator, svc blogservices.Service, fSvc fileservices.Service) Handler {
	return &defaultHandler{cv: cv, svc: svc, fileSvc: fSvc}
}
