package app

import (
	"cotizador_sounio_health/app/render"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/binding"
	"github.com/gobuffalo/envy"
)

var (
	root *buffalo.App
)

// App creates a new application with default settings and reading
// GO_ENV. It calls setRoutes to setup the routes for the app that's being
// created before returning it
func New() *buffalo.App {
	if root != nil {
		return root
	}

	root = buffalo.New(buffalo.Options{
		Env:         envy.Get("GO_ENV", "development"),
		SessionName: "_cotizador_sounio_health_session",
	})

	// Setting up libraries
	setupLibraries()
	// Setting the routes for the app
	setRoutes(root)

	return root
}

// setupLibraries allows to configure libraries that the app
// will use.
func setupLibraries() {
	binding.BaseRequestBinder = binding.NewRequestBinder(
		render.NewHTMLBinder(),
		binding.JSONContentTypeBinder{},
		binding.XMLRequestTypeBinder{},
	)

	binding.Register("application/html", render.HTMLBinding)
	binding.Register("text/html", render.HTMLBinding)
	binding.Register("application/x-www-form-urlencoded", render.HTMLBinding)
	binding.Register("html", render.HTMLBinding)
}
