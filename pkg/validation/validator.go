package validation

import (
	"github.com/go-playground/validator"
)

// Validator adapts the [validator.Validate] implementation to
// the [echo.Validator] interface.
type Validator struct {
	validator *validator.Validate
}

// Validate implements echo.Validator.
func (v *Validator) Validate(i interface{}) error {
    if err := v.validator.Struct(i); err != nil {
        return err
    }

    return nil
}

// NewValidator creates a new Validator.
func NewValidator() *Validator {
	return &Validator{validator.New()}
}
