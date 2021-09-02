package actions

import (
	"cotizador_sounio_health/app/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

func ClientsList(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	clients := models.Clients{}
	if err := tx.Order("first_name").All(&clients); err != nil {
		return err
	}

	c.Set("clients", clients)

	return c.Render(http.StatusOK, r.HTML("clients/index.plush.html"))
}

func NewClient(c buffalo.Context) error {
	c.Set("client", models.Client{})
	return c.Render(http.StatusOK, r.HTML("clients/new.plush.html"))
}

func CreateClient(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	client := models.Client{}

	if err := c.Bind(&client); err != nil {
		return err
	}

	if verrs := client.Validate(); verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("client", client)

		return c.Render(http.StatusOK, r.HTML("clients/new.plush.html"))
	}

	if err := tx.Create(&client); err != nil {
		return err
	}

	c.Flash().Add("success", "Cliente creado correctamente.")
	return c.Redirect(http.StatusSeeOther, "/clients")
}

// func ShowClient(c buffalo.Context) error {
// 	tx := c.Value("tx").(*pop.Connection)
// 	client := models.Client{}
// 	clientID := c.Param("client_id")

// 	if err := tx.Find(&client, clientID); err != nil {
// 		return err
// 	}

// 	c.Set("client", client)
// 	return c.Render(http.StatusOK, r.HTML("clients/show.plush.html"))
// }

func EditClient(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	client := models.Client{}
	clientID := c.Param("client_id")

	if err := tx.Find(&client, clientID); err != nil {
		return err
	}

	c.Set("client", client)
	return c.Render(http.StatusOK, r.HTML("clients/edit.plush.html"))
}

func UpdateClient(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	client := models.Client{}
	clientID := c.Param("client_id")

	if err := tx.Find(&client, clientID); err != nil {
		return err
	}

	if err := c.Bind(&client); err != nil {
		return err
	}

	if verrs := client.Validate(); verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("client", client)

		return c.Render(http.StatusOK, r.HTML("clients/edit.plush.html"))
	}

	if err := tx.Update(&client); err != nil {
		return err
	}

	c.Flash().Add("success", "Cliente modificado correctamente.")
	return c.Redirect(http.StatusSeeOther, "/clients")
}

func DeleteClient(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	client := models.Client{}
	clientID := c.Param("client_id")

	if err := tx.Find(&client, clientID); err != nil {
		return err
	}

	if err := tx.Destroy(&client); err != nil {
		return err
	}

	c.Flash().Add("success", "Cliente eliminado correctamente.")
	return c.Redirect(http.StatusSeeOther, "/clients")
}
