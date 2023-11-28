package data

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type RegisterFormData struct {
	Username             string `validate:"required,min=3,max=72"`
	Password             string `validate:"required,min=8,max=72"`
	PasswordConfirmation string `validate:"required,eqfield=Password"`
}

func (registerData *RegisterFormData) GatherFormErrors() (foundErrors bool, errors []string) {
	var validate = validator.New(validator.WithRequiredStructEnabled())
	var validationError = validate.Struct(registerData)
	errors = []string{}

	if validationError == nil {
		return false, errors
	}

	for _, err := range validationError.(validator.ValidationErrors) {
		errors = append(errors, messageFor(err))
	}

	return len(errors) != 0, errors
}

func messageFor(validationError validator.FieldError) string {
	switch validationError.Tag() {
	case "required":
		return fmt.Sprintf("%s is required.", validationError.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long.", validationError.Field(), validationError.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters long.", validationError.Field(), validationError.Param())
	case "eqfield":
		return "The passwords don't match."
	}
	return validationError.Error() // default error
}
