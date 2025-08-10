package utils

import (
	"database/sql"
	"errors"
	"net/http"
	"reflect"
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

func GenerateSelectColumns[T any](alias *string) string {
	var t T
	typ := reflect.TypeOf(t)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return ""
	}

	var columns []string
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if dbTag := field.Tag.Get("db"); dbTag != "" {
			if alias != nil {
				columns = append(columns, *alias+"."+dbTag)
			} else {
				columns = append(columns, dbTag)
			}
		}
	}

	return strings.Join(columns, ", ")
}

// ScanRowsToStructs maps sql.Rows to a slice of structs, handling NULLs
func ScanRowsToStructs[T any](rows *sql.Rows) ([]T, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []T
	colCount := len(columns)

	for rows.Next() {
		var t T
		dest := make([]interface{}, colCount)
		valuePtrs := make([]interface{}, colCount)

		val := reflect.ValueOf(&t).Elem()
		typ := val.Type()

		fieldMap := map[string]int{}
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			dbTag := field.Tag.Get("db")
			if dbTag != "" {
				fieldMap[dbTag] = i
			}
		}

		for i, col := range columns {
			if fieldIdx, ok := fieldMap[col]; ok {
				field := val.Field(fieldIdx)
				switch field.Kind() {
				case reflect.String:
					valuePtrs[i] = new(sql.NullString)
				case reflect.Int, reflect.Int64:
					valuePtrs[i] = new(sql.NullInt64)
				case reflect.Float64:
					valuePtrs[i] = new(sql.NullFloat64)
				case reflect.Bool:
					valuePtrs[i] = new(sql.NullBool)
				default:
					valuePtrs[i] = new(interface{})
				}
				dest[i] = valuePtrs[i]
			} else {
				var skip interface{}
				dest[i] = &skip
			}
		}

		if err := rows.Scan(dest...); err != nil {
			return nil, err
		}

		for i, col := range columns {
			if fieldIdx, ok := fieldMap[col]; ok {
				field := val.Field(fieldIdx)
				ptr := valuePtrs[i]

				switch v := ptr.(type) {
				case *sql.NullString:
					if v.Valid {
						field.SetString(v.String)
					}
				case *sql.NullInt64:
					if v.Valid {
						field.SetInt(v.Int64)
					}
				case *sql.NullFloat64:
					if v.Valid {
						field.SetFloat(v.Float64)
					}
				case *sql.NullBool:
					if v.Valid {
						field.SetBool(v.Bool)
					}
				case *interface{}:
					if field.CanSet() {
						field.Set(reflect.ValueOf(*v))
					}
				}
			}
		}

		results = append(results, t)
	}

	return results, nil
}

func ScanRowToStruct[T any](row *sql.Row, columns []string) (T, error) {
	var t T
	val := reflect.ValueOf(&t).Elem()
	typ := val.Type()

	if typ.Kind() != reflect.Struct {
		return t, errors.New("ScanRowToStruct expects a struct type")
	}

	fieldMap := make(map[string]int)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag != "" {
			fieldMap[dbTag] = i
		}
	}

	colCount := len(columns)
	dest := make([]interface{}, colCount)

	for i, col := range columns {
		if fieldIdx, ok := fieldMap[col]; ok {
			field := val.Field(fieldIdx)
			switch field.Kind() {
			case reflect.String:
				dest[i] = new(sql.NullString)
			case reflect.Int, reflect.Int64:
				dest[i] = new(sql.NullInt64)
			case reflect.Float64:
				dest[i] = new(sql.NullFloat64)
			case reflect.Bool:
				dest[i] = new(sql.NullBool)
			default:
				dest[i] = new(interface{})
			}
		} else {
			var skip interface{}
			dest[i] = &skip
		}
	}

	err := row.Scan(dest...)
	if err != nil {
		return t, err
	}

	for i, col := range columns {
		if fieldIdx, ok := fieldMap[col]; ok {
			field := val.Field(fieldIdx)
			ptr := dest[i]

			switch v := ptr.(type) {
			case *sql.NullString:
				if v.Valid {
					field.SetString(v.String)
				}
			case *sql.NullInt64:
				if v.Valid {
					field.SetInt(v.Int64)
				}
			case *sql.NullFloat64:
				if v.Valid {
					field.SetFloat(v.Float64)
				}
			case *sql.NullBool:
				if v.Valid {
					field.SetBool(v.Bool)
				}
			case *interface{}:
				if field.CanSet() {
					field.Set(reflect.ValueOf(*v))
				}
			}
		}
	}

	return t, nil
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
