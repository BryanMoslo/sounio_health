package actions

import (
	"cotizador_sounio_health/app/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v5"
)

func RateList(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	rates := models.InterestRates{}
	if err := tx.Order("contract_type, term, min_value").All(&rates); err != nil {
		return err
	}

	c.Set("interestRates", rates)

	return c.Render(http.StatusOK, r.HTML("interest_rates/index.plush.html"))
}

func NewInterestRate(c buffalo.Context) error {
	c.Set("interestRate", models.InterestRate{
		ContractType:      models.ContractTypeRent,
		PolicyRatePresent: true,
	})
	return c.Render(http.StatusOK, r.HTML("interest_rates/new.plush.html"))
}

func CreateInterestRate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	interestRate := models.InterestRate{}

	if err := c.Bind(&interestRate); err != nil {
		return err
	}

	if verrs := interestRate.Validate(); verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("interestRate", interestRate)

		return c.Render(http.StatusOK, r.HTML("interest_rates/new.plush.html"))
	}

	if interestRate.ContractType == models.ContractTypeRent {
		interestRate.PurchaseOptionPercentage = nulls.Float64{}
	}

	if !interestRate.PolicyRatePresent {
		interestRate.PolicyRate = nulls.Float64{}
	}

	if err := tx.Create(&interestRate); err != nil {
		return err
	}

	c.Flash().Add("success", "Tasa de intentés creada correctamente.")
	return c.Redirect(http.StatusSeeOther, "/interest_rates")
}

func ShowInterestRate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	interestRate := models.InterestRate{}
	interestRateID := c.Param("interest_rate_id")

	if err := tx.Find(&interestRate, interestRateID); err != nil {
		return err
	}

	c.Set("interestRate", interestRate)
	return c.Render(http.StatusOK, r.HTML("interest_rates/show.plush.html"))
}

func EditInterestRate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	interestRate := models.InterestRate{}
	interestRateID := c.Param("interest_rate_id")

	if err := tx.Find(&interestRate, interestRateID); err != nil {
		return err
	}

	c.Set("interestRate", interestRate)
	return c.Render(http.StatusOK, r.HTML("interest_rates/edit.plush.html"))
}

func UpdateInterestRate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	interestRate := models.InterestRate{}
	interestRateID := c.Param("interest_rate_id")

	if err := tx.Find(&interestRate, interestRateID); err != nil {
		return err
	}

	if err := c.Bind(&interestRate); err != nil {
		return err
	}

	if verrs := interestRate.Validate(); verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("interestRate", interestRate)

		return c.Render(http.StatusOK, r.HTML("interest_rates/edit.plush.html"))
	}

	if interestRate.ContractType == models.ContractTypeRent {
		interestRate.PurchaseOptionPercentage = nulls.Float64{}
	}

	if !interestRate.PolicyRatePresent {
		interestRate.PolicyRate = nulls.Float64{}
	}

	if err := tx.Update(&interestRate); err != nil {
		return err
	}

	c.Flash().Add("success", "Tasa de intentés modificada correctamente.")
	return c.Redirect(http.StatusSeeOther, "/interest_rates")
}

func DeleteInterestRate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	interestRate := models.InterestRate{}
	interestRateID := c.Param("interest_rate_id")

	if err := tx.Find(&interestRate, interestRateID); err != nil {
		return err
	}

	if err := tx.Destroy(&interestRate); err != nil {
		return err
	}

	c.Flash().Add("success", "Tasa de intentés eliminada correctamente.")
	return c.Redirect(http.StatusSeeOther, "/interest_rates")
}
