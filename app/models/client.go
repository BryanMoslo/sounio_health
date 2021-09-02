package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Client model struct.
type Client struct {
	ID          uuid.UUID `json:"id" db:"id"`
	FirstName   string    `json:"first_name" db:"first_name"`
	LastName    string    `json:"last_name" db:"last_name"`
	Email       string    `json:"email" db:"email"`
	Address     string    `json:"address" db:"address"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Clients array model struct of Client.
type Clients []Client

// String is not required by pop and may be deleted
func (c Client) String() string {
	jq, _ := json.Marshal(c)
	return string(jq)
}

// String is not required by pop and may be deleted
func (c Clients) String() string {
	jq, _ := json.Marshal(c)
	return string(jq)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Client) Validate() *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.FirstName, Name: "FirstName", Message: "Por favor, escriba el nombre del usuario."},
		&validators.StringIsPresent{Field: c.LastName, Name: "LastName", Message: "Por favor, escriba el apellido del usuario."},
		&validators.StringIsPresent{Field: c.Email, Name: "Email", Message: "Por favor, escriba el email del usuario."},
	)
}
