package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// Validator responsibility is to validate incoming request objects.
type Validator struct {
	validate *validator.Validate
}

// New instantiate new Validator.
func New() Validator {
	return Validator{validate: validator.New()}
}

// Validate incoming request object and returns error(s) if validation fails.
func (v Validator) Validate(data interface{}) error {
	// check if data is uuid and skip it
	if _, ok := data.(uuid.UUID); ok {
		return nil
	}
	return v.validate.Struct(data)
}
