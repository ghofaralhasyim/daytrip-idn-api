package utils

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateRequest(req interface{}, err error) ([]map[string]string, error) {
	var validationErrors []map[string]string

	val := reflect.TypeOf(req)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, errors.New("expected a struct type for validation")
	}

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrs {
			structField, _ := val.FieldByName(e.Field())
			jsonTag := structField.Tag.Get("json")
			if jsonTag == "" {
				jsonTag = strings.ToLower(e.Field())
			}

			friendlyMessage := GetFriendlyErrorMessage(e)
			validationErrors = append(validationErrors, map[string]string{
				jsonTag: friendlyMessage,
			})
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors, nil
	}
	return nil, nil
}

func GetFriendlyErrorMessage(e validator.FieldError) string {
	field := strings.ToLower(e.Field())
	switch e.Tag() {
	case "required":
		return field + " is required"
	case "email":
		return field + " must be a valid email address"
	case "min":
		return field + " must have at least " + e.Param() + " characters"
	case "max":
		return field + " must have no more than " + e.Param() + " characters"
	case "strongpassword":
		return field + " must be at least 8 characters long and include an uppercase letter, lowercase letter, number, and special character"
	default:
		return e.Error()
	}
}
