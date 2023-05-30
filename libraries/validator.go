package libraries

import (
	"github.com/go-playground/validator/v10"
	"net/mail"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.Validator.RegisterValidation("emailValidation", emailValidation)
	if err != nil {
		return err
	}

	err = cv.Validator.Struct(i)
	if err != nil {
		return err
	}
	return nil
}

func emailValidation(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != "" {
			_, err := mail.ParseAddress(value)
			if err != nil {
				return false
			}
		}
	}

	return true
}
