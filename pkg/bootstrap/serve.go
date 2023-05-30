package bootstrap

import (
	"github.com/godazz/bloGin/pkg/config"
	"github.com/godazz/bloGin/pkg/routing"
)

func Serve() {
	config.Set()

	routing.Init()

	routing.RegisterRoutes()

	routing.Serve()
}
