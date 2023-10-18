package models

import "github.com/go-playground/validator"

type Fact struct {
	ID       int    `json:"id,omitempty"`
	Question string `json:"question" validate:"required,min=1"`
	Answer   string `json:"answer" validate:"required,min=1"`
}

/* Start Region: Data validation */

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse

	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}

/* End Region: Data validation */

/* Ejemplos de uso del Validator. VER:
https://github.com/go-playground/validator/blob/master/_examples/simple/main.go
https://github.com/go-playground/validator/blob/master/_examples/struct-level/main.go
*/
