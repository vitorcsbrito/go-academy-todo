package todoerrors

import (
	"fmt"
)

type TaskNotFound struct {
	id string
}

func (e TaskNotFound) Error() string {
	return fmt.Sprintf("no task found for id %d", e.id)
}

func NewErrTaskNotFound(id string) error {
	return &TaskNotFound{id}
}

func NewErrResponse(id string) error {
	return &ErrorResponse{Message: id}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("oops, something happened: %s", e.Message)
}
