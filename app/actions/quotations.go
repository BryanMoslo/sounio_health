package actions

import (
	"cotizador_sounio_health/app/models"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

func QuotationsList(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	quotations := models.Quotations{}
	if err := tx.Eager().Order("updated_at").All(&quotations); err != nil {
		return err
	}

	c.Set("quotations", quotations)

	return c.Render(http.StatusOK, r.HTML("quotations/index.plush.html"))
}

func NewQuotation(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	clients := models.Clients{}
	if err := tx.Order("first_name").All(&clients); err != nil {
		return err
	}

	c.Set("quotation", models.Quotation{
		ContractType: models.ContractTypeRent,
		Term:         36,
	})

	clientsList := make(map[string]uuid.UUID)
	for _, client := range clients {
		clientsList[fmt.Sprintf("%v %v", client.FirstName, client.LastName)] = client.ID
	}

	c.Set("clients", clientsList)

	return c.Render(http.StatusOK, r.HTML("quotations/new.plush.html"))
}

func CreateQuotation(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	quotation := models.Quotation{}

	if err := c.Bind(&quotation); err != nil {
		return err
	}

	if verrs := quotation.Validate(); verrs.HasAny() {
		clients := models.Clients{}
		if err := tx.Order("first_name").All(&clients); err != nil {
			return err
		}

		clientsList := make(map[string]uuid.UUID)
		for _, client := range clients {
			clientsList[fmt.Sprintf("%v %v", client.FirstName, client.LastName)] = client.ID
		}

		c.Set("clients", clientsList)
		c.Set("errors", verrs)
		c.Set("quotation", quotation)

		return c.Render(http.StatusOK, r.HTML("quotations/new.plush.html"))
	}

	rate := models.InterestRate{}
	err := tx.Where(
		"contract_type = ?", quotation.ContractType).Where(
		"term = ?", quotation.Term).Where(
		"min_value <= ?", quotation.EquipmentValue).Where(
		"max_value >= ?", quotation.EquipmentValue).First(&rate)

	if err != nil && errors.Cause(err) != sql.ErrNoRows {
		return err
	}

	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		clients := models.Clients{}
		if err := tx.Order("first_name").All(&clients); err != nil {
			return err
		}

		clientsList := make(map[string]uuid.UUID)
		for _, client := range clients {
			clientsList[fmt.Sprintf("%v %v", client.FirstName, client.LastName)] = client.ID
		}

		c.Set("clients", clientsList)
		c.Set("quotation", quotation)

		c.Flash().Add("danger", fmt.Sprintf("No existe una tasa de interés en el sistema para un %v a %v meses por un equipo de %v.", quotation.ContractType, quotation.Term, quotation.FormatEquipmentValue()))
		return c.Render(http.StatusOK, r.HTML("quotations/new.plush.html"))
	}

	quotation.RateID = rate.ID
	quotation.CalculateFee(rate)

	if quotation.ContractType == models.ContractTypeLeasing {
		quotation.CalculatePurchaseOptionValue(rate)
	}

	if err := tx.Create(&quotation); err != nil {
		return err
	}

	c.Flash().Add("success", "Cotización creada correctamente.")
	return c.Redirect(http.StatusSeeOther, "/quotations/%v/show", quotation.ID)
}

func ShowQuotation(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	quotation := models.Quotation{}
	quotationID := c.Param("quotation_id")

	if err := tx.Eager().Find(&quotation, quotationID); err != nil {
		return err
	}

	c.Set("quotation", quotation)
	return c.Render(http.StatusOK, r.HTML("quotations/show.plush.html"))
}

func EditQuotation(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	clients := models.Clients{}
	quotation := models.Quotation{}
	quotationID := c.Param("quotation_id")

	if err := tx.Find(&quotation, quotationID); err != nil {
		return err
	}

	if err := tx.Order("first_name").All(&clients); err != nil {
		return err
	}

	clientsList := make(map[string]uuid.UUID)
	for _, client := range clients {
		clientsList[fmt.Sprintf("%v %v", client.FirstName, client.LastName)] = client.ID
	}

	c.Set("clients", clientsList)
	c.Set("quotation", quotation)
	return c.Render(http.StatusOK, r.HTML("quotations/edit.plush.html"))
}

func UpdateQuotation(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	quotation := models.Quotation{}
	quotationID := c.Param("quotation_id")

	if err := tx.Find(&quotation, quotationID); err != nil {
		return err
	}

	if err := c.Bind(&quotation); err != nil {
		return err
	}

	if verrs := quotation.Validate(); verrs.HasAny() {
		clients := models.Clients{}
		if err := tx.Order("first_name").All(&clients); err != nil {
			return err
		}

		clientsList := make(map[string]uuid.UUID)
		for _, client := range clients {
			clientsList[fmt.Sprintf("%v %v", client.FirstName, client.LastName)] = client.ID
		}

		c.Set("clients", clientsList)
		c.Set("errors", verrs)
		c.Set("quotation", quotation)

		return c.Render(http.StatusOK, r.HTML("quotations/edit.plush.html"))
	}

	rate := models.InterestRate{}
	err := tx.Where(
		"contract_type = ?", quotation.ContractType).Where(
		"term = ?", quotation.Term).Where(
		"min_value <= ?", quotation.EquipmentValue).Where(
		"max_value >= ?", quotation.EquipmentValue).First(&rate)

	if err != nil && errors.Cause(err) != sql.ErrNoRows {
		return err
	}

	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		clients := models.Clients{}
		if err := tx.Order("first_name").All(&clients); err != nil {
			return err
		}

		clientsList := make(map[string]uuid.UUID)
		for _, client := range clients {
			clientsList[fmt.Sprintf("%v %v", client.FirstName, client.LastName)] = client.ID
		}

		c.Set("clients", clientsList)
		c.Set("quotation", quotation)

		c.Flash().Add("danger", fmt.Sprintf("No existe una tasa de interés en el sistema para un %v a %v meses por un equipo de %v.", quotation.ContractType, quotation.Term, quotation.FormatEquipmentValue()))
		return c.Render(http.StatusOK, r.HTML("quotations/new.plush.html"))
	}

	quotation.RateID = rate.ID
	quotation.CalculateFee(rate)

	if quotation.ContractType == models.ContractTypeLeasing {
		quotation.CalculatePurchaseOptionValue(rate)
	}

	if err := tx.Update(&quotation); err != nil {
		return err
	}

	c.Flash().Add("success", "Cotización modificada correctamente.")
	return c.Redirect(http.StatusSeeOther, "/quotations")
}

func DeleteQuotation(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	quotation := models.Quotation{}
	quotationID := c.Param("quotation_id")

	if err := tx.Find(&quotation, quotationID); err != nil {
		return err
	}

	if err := tx.Destroy(&quotation); err != nil {
		return err
	}

	c.Flash().Add("success", "Cotización eliminada correctamente.")
	return c.Redirect(http.StatusSeeOther, "/quotations")
}
