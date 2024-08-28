package task

import (
	"github.com/google/uuid"
	. "github.com/vitorcsbrito/go-academy-todo/model"
)

type Repository interface {
	SaveTask(task Task) (uuid.UUID, error)
	UpdateTask(id uuid.UUID, task Task) (uuid.UUID, error)
	FindById(id uuid.UUID) (*Task, uuid.UUID, error)
	DeleteTask(taskId *Task) error
	FindAllTasks() ([]Task, error)
}
