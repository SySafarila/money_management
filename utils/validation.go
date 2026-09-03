package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
}

func ValidateVariable(data interface{}, varIdentifier string, validation string) error {
	errString := make(map[string]string)
	errs := validate.Var(data, validation)
	if errs != nil {
		var validateErrs validator.ValidationErrors
		if errors.As(errs, &validateErrs) {
			for _, e := range validateErrs {
				errString[varIdentifier] = fmt.Sprintf("%s is not a valid %s", e.Value(), e.Tag())
			}
		}
		return &ValidationError{errString}
	}
	return nil
}

func ValidateStruct(data interface{}) error {
	err := validate.Struct(data)
	if err != nil {
		validationErrors := formatValidationErrors(err, data)
		return &ValidationError{validationErrors}
	}
	return nil
}

func formatValidationErrors(err error, data interface{}) map[string]string {
	errorsMap := make(map[string]string)

	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		return errorsMap
	}

	t := reflect.TypeOf(data)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for _, fieldErr := range validationErrors {
		field := fieldErr.StructField()

		structField, ok := t.FieldByName(field)
		if !ok {
			continue
		}

		jsonField := strings.Split(
			structField.Tag.Get("json"),
			",",
		)[0]

		if jsonField == "" || jsonField == "-" {
			jsonField = field
		}

		errorsMap[jsonField] = getValidationMessage(
			jsonField,
			fieldErr.Tag(),
			fieldErr.Param(),
		)
	}

	return errorsMap
}
func getValidationMessage(field, tag, param string) string {
	fieldName := strings.ToLower(field)

	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", fieldName)

	case "email":
		return fmt.Sprintf("%s must be a valid email address", fieldName)

	case "min":
		return fmt.Sprintf("%s must be at least %s characters", fieldName, param)

	case "max":
		return fmt.Sprintf("%s must be at most %s characters", fieldName, param)

	case "len":
		return fmt.Sprintf("%s must be exactly %s characters", fieldName, param)

	case "numeric":
		return fmt.Sprintf("%s must be a number", fieldName)

	case "oneof":
		return fmt.Sprintf("%s must be one of: %s", fieldName, param)

	case "url":
		return fmt.Sprintf("%s must be a valid URL", fieldName)

	default:
		return fmt.Sprintf("%s is invalid", fieldName)
	}
}
