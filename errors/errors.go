package errors

import "fmt"

type TaskNotFound struct {
	id int
}

func (e TaskNotFound) Error() string {
	return fmt.Sprintf("no task found for id %d", e.id)
}

func NewErrTaskNotFound(id int) error {
	return &TaskNotFound{id: id}
}

type ErrorResponse struct {
	Message string `json:"message"`
}
