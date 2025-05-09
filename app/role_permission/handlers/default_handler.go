package rphandlers

import (
	rpservices "github.com/meth-suchatchai/kz-blog-api/app/role_permission/services"
	"github.com/meth-suchatchai/kz-blog-api/lib/validator"
)

type defaultHandler struct {
	cv  validator.CustomValidator
	svc rpservices.Service
}

func NewHandler(cv validator.CustomValidator, svc rpservices.Service) Handler {
	return &defaultHandler{cv: cv, svc: svc}
}
