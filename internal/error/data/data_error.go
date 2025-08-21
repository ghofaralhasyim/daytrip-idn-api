package error_data

import (
	"fmt"
	"net/http"
)

const (
	InvalidDataRequest = "INVALID_DATA_REQUEST"
	UnknownError       = "UNKNOWN_ERROR"
)

type ErrorDetail struct {
	Message    string
	StatusCode int
}

var Details = map[string]ErrorDetail{
	InvalidDataRequest: {Message: "failed to get data from repository", StatusCode: http.StatusBadRequest},
	UnknownError:       {Message: "unknown error occurred", StatusCode: http.StatusInternalServerError},
}

type AppError struct {
	Code       int
	Message    string
	StatusCode int
	Err        error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%d] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func NewAppError(code string, err error) *AppError {
	detail, ok := Details[code]
	if !ok {
		detail = Details[UnknownError]
	}
	return &AppError{
		Code:       detail.StatusCode,
		Message:    detail.Message,
		StatusCode: detail.StatusCode,
		Err:        err,
	}
}
