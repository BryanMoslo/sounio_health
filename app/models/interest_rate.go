package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

var (
	ContractTypeRent    = "Arriendo"
	ContractTypeLeasing = "Leasing"
)

// InterestRate model struct.
type InterestRate struct {
	ID                       uuid.UUID     `json:"id" db:"id"`
	ContractType             string        `json:"contract_type" db:"contract_type"`
	Rate                     nulls.Float64 `json:"rate" db:"rate"`
	Term                     int64         `json:"term" db:"term"`
	PolicyRatePresent        bool          `json:"policy_rate_present" db:"policy_rate_present"`
	PolicyRate               nulls.Float64 `json:"policy_rate" db:"policy_rate"`
	MinValue                 nulls.Float64 `json:"min_value" db:"min_value"`
	MaxValue                 nulls.Float64 `json:"max_value" db:"max_value"`
	PurchaseOptionPercentage nulls.Float64 `json:"purchase_option_percentage" db:"purchase_option_percentage"`
	CreatedAt                time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt                time.Time     `json:"updated_at" db:"updated_at"`
}

// InterestRates array model struct of InterestRate.
type InterestRates []InterestRate

// String is not required by pop and may be deleted
func (c InterestRate) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// String is not required by pop and may be deleted
func (c InterestRates) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (ir *InterestRate) Validate() *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: ir.ContractType, Name: "ContractType", Message: "Por favor, selecciona un tipo de contrato."},
		&validators.IntIsPresent{Field: int(ir.Term), Name: "Term", Message: "Por favor, selecciona un plazo."},
		&validators.FuncValidator{
			Name:    "Rate",
			Message: "%v Por favor, ingrese una tasa de interés.",
			Fn: func() bool {
				return ir.Rate.Valid
			},
		},
		&validators.FuncValidator{
			Name:    "Rate",
			Message: "%v Por favor, ingrese una tasa de interés valida.",
			Fn: func() bool {
				return !ir.Rate.Valid || ir.Rate.Float64 > 0
			},
		},
		&validators.FuncValidator{
			Name:    "PolicyRate",
			Message: "%v Por favor, ingrese una tasa de interés de póliza.",
			Fn: func() bool {
				if !ir.PolicyRatePresent {
					return true
				}

				return ir.PolicyRate.Valid
			},
		},
		&validators.FuncValidator{
			Name:    "PolicyRate",
			Message: "%v Por favor, ingrese una tasa de interés de póliza valida.",
			Fn: func() bool {
				if !ir.PolicyRatePresent {
					return true
				}

				return !ir.PolicyRate.Valid || ir.PolicyRate.Float64 > 0
			},
		},
		&validators.FuncValidator{
			Name:    "MaxValue",
			Message: "%v Por favor, ingrese un  valor máximo.",
			Fn: func() bool {
				return ir.MaxValue.Valid
			},
		},
		&validators.FuncValidator{
			Name:    "MaxValue",
			Message: "%v Por favor, ingrese un valor máximo.",
			Fn: func() bool {
				return !ir.MaxValue.Valid || ir.MaxValue.Float64 > 0
			},
		},
		&validators.FuncValidator{
			Name:    "MaxValue",
			Message: "%v El valor máximo debe ser mayor al mínimo.",
			Fn: func() bool {
				if ir.MinValue.Valid && ir.MaxValue.Valid && ir.MaxValue.Float64 > 0 {
					return ir.MaxValue.Float64 > ir.MinValue.Float64
				}

				return true
			},
		},
		&validators.FuncValidator{
			Name:    "PurchaseOptionPercentage",
			Message: "%v Por favor, ingrese el porcentaje de opción de compra.",
			Fn: func() bool {
				if ir.ContractType == ContractTypeLeasing {
					return ir.PurchaseOptionPercentage.Valid && ir.PurchaseOptionPercentage.Float64 > 0
				}

				return true
			},
		},
	)
}
