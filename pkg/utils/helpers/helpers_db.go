package helpers

import (
	"database/sql"
	"errors"
	"reflect"
	"strings"
	"time"
)

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

func NullString(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

func NullInt64(ni sql.NullInt64) *int64 {
	if ni.Valid {
		return &ni.Int64
	}
	return nil
}

func NullTime(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}

func NewNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: *s, Valid: true}
}

func NewNullInt64(i *int64) sql.NullInt64 {
	if i == nil {
		return sql.NullInt64{Valid: false}
	}
	return sql.NullInt64{Int64: *i, Valid: true}
}

func NewNullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: *t, Valid: true}
}
