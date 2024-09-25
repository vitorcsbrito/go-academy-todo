package errors

import (
	"errors"
)

var (
	ErrTaskIdMissingFromRequest = errors.New("task id missing from request")
	ErrTaskNotFound             = errors.New("task not found")
	ErrTaskDescriptionNotFound  = errors.New("missing task description")
	ErrMissingErrorDetails      = errors.New("missing user details")
	ErrEmailTaken               = errors.New("email unavailable")
	ErrInvalidToken             = errors.New("invalid token")
	ErrExpiredToken             = errors.New("expired token")
	ErrMissingAuthHeader        = errors.New("missing authorization header")
	ErrNoUsernameFound          = errors.New("no username found")
	ErrInvalidCredentials       = errors.New("invalid credentials")
	ErrUserNotFound             = errors.New("user not found")
)

func NewErrResponse(err error) error {
	return &ErrorResponse{Message: err.Error()}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}
