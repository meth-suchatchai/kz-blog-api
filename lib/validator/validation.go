package validator

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	Validator *validator.Validate
}

type ErrorResponse struct {
	Field string
	Value interface{}
}

var Validate = validator.New()

func (cv CustomValidator) Validate(data any) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := Validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.Field = err.Field() // Export struct field name
			elem.Value = err.Value() // Export field value

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
