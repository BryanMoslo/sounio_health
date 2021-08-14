package app

import (
	base "cotizador_sounio_health"
	"cotizador_sounio_health/app/actions/home"
	"cotizador_sounio_health/app/middleware"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(root *buffalo.App) {
	root.Use(middleware.Transaction)
	root.Use(middleware.ParameterLogger)
	root.Use(middleware.CSRF)

	root.GET("/", home.Index)
	root.ServeFiles("/", base.Assets)
}