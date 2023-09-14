package clienthandlers

import (
	blogservices "github.com/kuroshibaz/app/blog/services"
	clientservices "github.com/kuroshibaz/app/client/services"
	etcdservices "github.com/kuroshibaz/app/etcd/services"
	"github.com/kuroshibaz/app/user/services"
	kzjwt "github.com/kuroshibaz/lib/jwt"
	"github.com/kuroshibaz/lib/validator"
)

type defaultHandler struct {
	validate      validator.CustomValidator
	userService   userservices.Service
	etcdService   etcdservices.Service
	clientService clientservices.Service
	blogService   blogservices.Service
	auth          kzjwt.AuthJWT
}

func NewHandler(validate validator.CustomValidator, userService userservices.Service, blogService blogservices.Service, etcdService etcdservices.Service, clientService clientservices.Service, auth kzjwt.AuthJWT) Handler {
	return &defaultHandler{validate: validate, userService: userService, blogService: blogService, etcdService: etcdService, clientService: clientService, auth: auth}
}
