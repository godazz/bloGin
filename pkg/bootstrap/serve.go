package bootstrap

import (
	"github.com/godazz/bloGin/pkg/config"
	"github.com/godazz/bloGin/pkg/html"
	"github.com/godazz/bloGin/pkg/routing"
	"github.com/godazz/bloGin/pkg/static"
)

func Serve() {
	config.Set()

	routing.Init()

	static.LoadStatic(routing.GetRouter())

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
