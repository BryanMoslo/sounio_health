package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func Home(c buffalo.Context) error {
	return c.Redirect(http.StatusSeeOther, "/quotations")

}
