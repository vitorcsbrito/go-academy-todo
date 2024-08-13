package repository

import (
	"github.com/google/uuid"
	. "go-todo-app/errors"
	. "go-todo-app/model"
)

type TaskRepository interface {
	SaveTask(task Task) uuid.UUID
	UpdateTask(id uuid.UUID, task Task) (uuid.UUID, error)
	FindTaskById(id uuid.UUID) (*Task, uuid.UUID, error)
	DeleteTask(taskId *Task) error
	FindAllTasks() ([]Task, error)
}

func (s *Repository) SaveTask(task Task) uuid.UUID {

	db := GetInstance().DB

	t := Task{Description: task.Description, Done: task.Done}

	newUUID, _ := uuid.NewUUID()
	t.Id = newUUID

	_ = db.Save(&t)

	return t.Id
}

func (s *Repository) UpdateTask(id uuid.UUID, task Task) (i uuid.UUID, err error) {
	t, i, err := s.FindTaskById(id)

	if err != nil {
		return i, err
	}

	t.Description = task.Description
	t.Done = task.Done

	res := s.DB.Save(&t)

	if res.Error != nil {
		return i, res.Error
	}

	return i, nil
}

func (s *Repository) FindTaskById(id uuid.UUID) (*Task, uuid.UUID, error) {

	var foundTask Task
	res := s.DB.Find(&foundTask, id)

	newUUID, _ := uuid.NewUUID()
	if res.Error == nil {
		return &foundTask, newUUID, nil
	}

	return nil, newUUID, NewErrTaskNotFound(id)
}

func (s *Repository) FindAllTasks() ([]Task, error) {

	var foundTask []Task

	res := s.DB.Order("created_at asc").Find(&foundTask)
	if res.Error != nil {
		return nil, res.Error
	}

	return foundTask, nil
}

func (s *Repository) DeleteTask(task *Task) error {
	res := s.DB.Delete(task)

	return res.Error
}
