package router

import (
	"github.com/meth-suchatchai/kuroctxfiber"
)

func NewKuroRouter(opts *Options) kuroctxfiber.KuroFiber {
	app := kuroctxfiber.New(kuroctxfiber.Config{Name: "kz-blog-api"})

	return app
}
