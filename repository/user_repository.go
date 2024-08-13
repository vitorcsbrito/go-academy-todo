package repository

import (
	"fmt"
	"github.com/google/uuid"
	. "go-todo-app/model"
)

type UserRepository interface {
	SaveUser(User) uuid.UUID
	//UpdateTask(id uuid.UUID, task Task) (uuid.UUID, error)
	//FindTaskById(id uuid.UUID) (*Task, uuid.UUID, error)
	//DeleteTask(taskId *Task) error
	//FindAllTasks() ([]Task, error)
}

func (s *Repository) SaveUser(user User) uuid.UUID {

	_ = GetInstance().DB

	newUUID, _ := uuid.NewUUID()
	user.Id = newUUID

	res := GetInstance().DB.Save(&user)

	if res.Error != nil {
		fmt.Printf("error saving user: %v", res.Error)
	}

	return newUUID
}

//func (s *Repository) UpdateTask(id uuid.UUID, task Task) (i uuid.UUID, err error) {
//	t, i, err := s.FindTaskById(id)
//
//	if err != nil {
//		return i, err
//	}
//
//	t.Description = task.Description
//	t.Done = task.Done
//
//	res := s.DB.SaveTask(&t)
//
//	if res.Error != nil {
//		return i, res.Error
//	}
//
//	return i, nil
//}
//
//func (s *Repository) FindTaskById(id uuid.UUID) (*Task, uuid.UUID, error) {
//
//	var foundTask Task
//	res := s.DB.Find(&foundTask, id)
//
//	newUUID, _ := uuid.NewUUID()
//	if res.Error == nil {
//		return &foundTask, newUUID, nil
//	}
//
//	return nil, newUUID, NewErrTaskNotFound(id)
//}
//
//func (s *Repository) FindAllTasks() ([]Task, error) {
//
//	var foundTask []Task
//
//	res := s.DB.Order("created_at asc").Find(&foundTask)
//	if res.Error != nil {
//		return nil, res.Error
//	}
//
//	return foundTask, nil
//}
//
//func (s *Repository) DeleteTask(task *Task) error {
//	res := s.DB.DeleteTask(task)
//
//	return res.Error
//}
