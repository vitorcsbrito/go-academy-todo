package mysql

import (
	"github.com/google/uuid"
	. "github.com/vitorcsbrito/go-academy-todo/model"
	. "github.com/vitorcsbrito/go-academy-todo/repository"
)

type MySqlRepository struct {
	*Repository
}

func NewMySqlRepository(repository *Repository) *MySqlRepository {
	repository.Init(GetMySQLConnection())

	return &MySqlRepository{repository}
}

func (s *MySqlRepository) SaveTask(task Task) (uuid.UUID, error) {
	newUUID, _ := uuid.NewUUID()
	task.ID = newUUID

	res := s.DB.Create(&task)

	return task.ID, res.Error
}

func (s *MySqlRepository) UpdateTask(id uuid.UUID, task Task) (i uuid.UUID, err error) {
	t, i, err := s.FindById(id)

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

func (s *MySqlRepository) FindById(id uuid.UUID) (*Task, uuid.UUID, error) {
	var foundTask Task
	res := s.DB.First(&foundTask, id)

	return &foundTask, id, res.Error
}

func (s *MySqlRepository) FindAllTasks() ([]Task, error) {

	var tasks []Task

	res := s.DB.Order("created_at asc").Find(&tasks)

	return tasks, res.Error
}

func (s *MySqlRepository) DeleteTask(task *Task) error {
	res := s.DB.Delete(task)

	return res.Error
}
