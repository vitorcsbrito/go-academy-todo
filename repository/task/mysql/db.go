package mysql

import (
	"github.com/google/uuid"
	. "github.com/vitorcsbrito/go-academy-todo/model/task"
	. "github.com/vitorcsbrito/go-academy-todo/repository"
	. "github.com/vitorcsbrito/utils/errors"
)

type MySqlRepository struct {
	*Repository
}

func NewMySqlRepository(repository *Repository) *MySqlRepository {
	repository.Init(GetMySQLConnection())

	return &MySqlRepository{repository}
}

func (s *MySqlRepository) SaveTask(task Task) (uuid.UUID, error) {
	t := Task{Description: task.Description, Done: task.Done}

	newUUID, _ := uuid.NewUUID()
	t.Id = newUUID

	res := s.DB.Create(&t)

	if res.Error != nil {
		return newUUID, res.Error
	}

	return t.Id, nil
}

func (s *MySqlRepository) UpdateTask(id uuid.UUID, task Task) (i uuid.UUID, err error) {
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

func (s *MySqlRepository) FindTaskById(id uuid.UUID) (*Task, uuid.UUID, error) {
	var foundTask Task
	res := s.DB.Find(&foundTask, id)

	newUUID, _ := uuid.NewUUID()
	if res.Error == nil && res.RowsAffected > 0 {
		return &foundTask, newUUID, nil
	}

	return nil, newUUID, ErrTaskNotFound
}

func (s *MySqlRepository) FindAllTasks() ([]Task, error) {

	var foundTask []Task

	res := s.DB.Order("created_at asc").Find(&foundTask)
	if res.Error != nil {
		return nil, res.Error
	}

	return foundTask, nil
}

func (s *MySqlRepository) DeleteTask(task *Task) error {
	res := s.DB.Delete(task)

	return res.Error
}
