package error_app

import (
	"fmt"
	"net/http"
)

const (
	RepositoryGetError    = "REPOSITORY_GET_ERROR"
	RepositorySaveError   = "REPOSITORY_SAVE_ERROR"
	RepositoryDeleteError = "REPOSITORY_DELETE_ERROR"
	UsecaseValidateError  = "USECASE_VALIDATE_ERROR"
	InvalidCredentials    = "INVALID_CREDENTIALS"
	UnknownError          = "UNKNOWN_ERROR"
)

type ErrorDetail struct {
	Message    string
	StatusCode int
}

var Details = map[string]ErrorDetail{
	RepositoryGetError:    {Message: "failed to get data from repository", StatusCode: http.StatusInternalServerError},
	RepositorySaveError:   {Message: "failed to save data to repository", StatusCode: http.StatusInternalServerError},
	RepositoryDeleteError: {Message: "failed to delete data from repository", StatusCode: http.StatusInternalServerError},
	UsecaseValidateError:  {Message: "failed to validate request data", StatusCode: http.StatusBadRequest},
	InvalidCredentials:    {Message: "invalid credentials", StatusCode: http.StatusForbidden},
	UnknownError:          {Message: "unknown error occurred", StatusCode: http.StatusInternalServerError},
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
