package errors

import (
	"fmt"
	"github.com/google/uuid"
)

type TaskNotFound struct {
	id uuid.UUID
}

func (e TaskNotFound) Error() string {
	return fmt.Sprintf("no task found for id %d", e.id)
}

func NewErrTaskNotFound(id uuid.UUID) error {
	return &TaskNotFound{id: id}
}

type ErrorResponse struct {
	Message string `json:"message"`
}
