package clienthandlers

import (
	blogservices "github.com/meth-suchatchai/kz-blog-api/app/blog/services"
	clientservices "github.com/meth-suchatchai/kz-blog-api/app/client/services"
	etcdservices "github.com/meth-suchatchai/kz-blog-api/app/etcd/services"
	"github.com/meth-suchatchai/kz-blog-api/app/user/services"
	kzjwt "github.com/meth-suchatchai/kz-blog-api/lib/jwt"
	"github.com/meth-suchatchai/kz-blog-api/lib/validator"
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
