package app

import (
	base "cotizador_sounio_health"
	"cotizador_sounio_health/app/actions"
	"cotizador_sounio_health/app/middleware"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(root *buffalo.App) {
	root.Use(middleware.Transaction)
	root.Use(middleware.ParameterLogger)
	root.Use(middleware.CSRF)

	root.GET("/interest_rates", actions.RateList)
	root.GET("/interest_rates/new", actions.NewInterestRate)
	root.POST("/interest_rates/create", actions.CreateInterestRate)
	root.GET("/interest_rates/{interest_rate_id}/edit", actions.EditInterestRate)
	root.PUT("/interest_rates/{interest_rate_id}/update", actions.UpdateInterestRate)
	root.DELETE("/interest_rates/{interest_rate_id}/delete", actions.DeleteInterestRate)

	root.GET("/quotations", actions.QuotationsList)
	root.GET("/quotations/new", actions.NewQuotation)
	root.POST("/quotations/create", actions.CreateQuotation)
	root.GET("/quotations/{quotation_id}/show", actions.ShowQuotation)
	root.GET("/quotations/{quotation_id}/edit", actions.EditQuotation)
	root.PUT("/quotations/{quotation_id}/update", actions.UpdateQuotation)
	root.DELETE("/quotations/{quotation_id}/delete", actions.DeleteQuotation)

	root.GET("/clients", actions.ClientsList)
	root.GET("/clients/new", actions.NewClient)
	root.POST("/clients/create", actions.CreateClient)
	// root.GET("/clients/{client_id}/show", actions.ShowClient)
	root.GET("/clients/{client_id}/edit", actions.EditClient)
	root.PUT("/clients/{client_id}/update", actions.UpdateClient)
	root.DELETE("/clients/{client_id}/delete", actions.DeleteClient)

	root.ServeFiles("/", base.Assets)
}
