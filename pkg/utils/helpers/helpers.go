package helpers

import (
	"errors"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	error_app "github.com/daytrip-idn-api/internal/error"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
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

func EchoError(ctx echo.Context, err error) error {
	if err == nil {
		return nil
	}

	if appErr, ok := err.(*error_app.AppError); ok {
		return ctx.JSON(appErr.StatusCode, map[string]interface{}{
			"code":    appErr.Code,
			"message": appErr.Message,
		})
	}

	return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
		"code":    "UNKNOWN_ERROR",
		"message": "something went wrong",
	})
}

func GenerateSlug(title string) string {
	// Lowercase
	slug := strings.ToLower(title)

	// Replace spaces with -
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove non-alphanumeric or dash
	reg := regexp.MustCompile(`[^a-z0-9-]+`)
	slug = reg.ReplaceAllString(slug, "")

	// Trim trailing dashes
	slug = strings.Trim(slug, "-")

	return slug
}
