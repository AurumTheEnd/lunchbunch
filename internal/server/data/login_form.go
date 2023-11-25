package data

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type LoginFormData struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

func (loginData *LoginFormData) GatherFormErrors() (foundErrors bool, errors []string) {
	var validate = validator.New(validator.WithRequiredStructEnabled())
	var validationError = validate.Struct(loginData)
	errors = []string{}

	if validationError == nil {
		return false, errors
	}

	for _, err := range validationError.(validator.ValidationErrors) {
		errors = append(errors, loginData.messageFor(err))
	}

	return len(errors) != 0, errors
}

func (_ *LoginFormData) messageFor(validationError validator.FieldError) string {
	switch validationError.Tag() {
	case "required":
		return fmt.Sprintf("%s is required.", validationError.Field())
	}
	return validationError.Error() // default error
}
