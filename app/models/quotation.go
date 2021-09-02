package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Quotation model struct.
type Quotation struct {
	ID                   uuid.UUID     `json:"id" db:"id"`
	ContractType         string        `json:"contract_type" db:"contract_type"`
	EquipmentValue       nulls.Float64 `json:"equipment_value" db:"equipment_value"`
	EquipmentName        string        `json:"equipment_name" db:"equipment_name"`
	EquipmentDescription string        `json:"equipment_description" db:"equipment_description"`
	Term                 int64         `json:"term" db:"term"`
	RateID               uuid.UUID     `json:"rate_id" db:"rate_id"`
	Rate                 InterestRate  `belongs_to:"interest_rates" fk_id:"RateID" db:"-"`
	ClientID             nulls.UUID    `json:"client_id" db:"client_id"`
	Client               Client        `belongs_to:"clients" fk_id:"ClientID" db:"-"`
	Fee                  float64       `json:"fee" db:"fee"`
	CreatedAt            time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time     `json:"updated_at" db:"updated_at"`
}

// Quotations array model struct of Quotation.
type Quotations []Quotation

// String is not required by pop and may be deleted
func (q Quotation) String() string {
	jq, _ := json.Marshal(q)
	return string(jq)
}

// String is not required by pop and may be deleted
func (q Quotations) String() string {
	jq, _ := json.Marshal(q)
	return string(jq)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (q *Quotation) Validate() *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: q.ContractType, Name: "ContractType", Message: "Por favor, selecciona un tipo de contrato."},
		&validators.StringIsPresent{Field: q.EquipmentName, Name: "EquipmentName", Message: "Por favor, escriba el nombre el equipo."},
		&validators.IntIsPresent{Field: int(q.Term), Name: "Term", Message: "Por favor, selecciona un plazo."},
		&validators.FuncValidator{
			Name:    "EquipmentValue",
			Message: "%v Por favor, ingrese el valor del equipo.",
			Fn: func() bool {
				return q.EquipmentValue.Valid
			},
		},
		&validators.FuncValidator{
			Name:    "EquipmentValue",
			Message: "%v Por favor, ingrese un valor de equipo valido.",
			Fn: func() bool {
				return !q.EquipmentValue.Valid || q.EquipmentValue.Float64 > 0
			},
		},
	)
}

func (q *Quotation) CalculateFee(rate InterestRate) {
	totalRate := rate.Rate.Float64
	if rate.PolicyRatePresent {
		totalRate += rate.PolicyRate.Float64
	}

	fee := q.EquipmentValue.Float64 * totalRate / 100
	q.Fee = fee
}
