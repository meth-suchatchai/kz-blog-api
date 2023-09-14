package rphandlers

import (
	rpservices "github.com/kuroshibaz/app/role_permission/services"
	"github.com/kuroshibaz/lib/validator"
)

type defaultHandler struct {
	cv  validator.CustomValidator
	svc rpservices.Service
}

func NewHandler(cv validator.CustomValidator, svc rpservices.Service) Handler {
	return &defaultHandler{cv: cv, svc: svc}
}
