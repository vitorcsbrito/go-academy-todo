package errors

import (
	"errors"
	"fmt"
)

var (
	ErrTaskIdMissingFromRequest = errors.New("task id missing from request")
	ErrTaskNotFound             = errors.New("task not found")
	ErrTaskDescriptionNotFound  = errors.New("missing task description")
	ErrMissingErrorDetails      = errors.New("missing user details")
	ErrEmailTaken               = errors.New("email unavailable")
)

func NewErrResponse(err error) error {
	return &ErrorResponse{Message: err.Error()}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("oops, something happened: %s", e.Message)
}
