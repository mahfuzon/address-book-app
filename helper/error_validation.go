package helper

import "github.com/go-playground/validator/v10"

func ConverseToErrorString(errors validator.ValidationErrors) []string {
	var errorMessage []string

	for _, fieldError := range errors {
		errorMessage = append(errorMessage, fieldError.Error())
	}

	return errorMessage
}
