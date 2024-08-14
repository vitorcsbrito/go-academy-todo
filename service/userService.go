package service

import (
	"github.com/google/uuid"
	. "model"
	. "repository"
)

type UserService struct {
	userRepository UserRepository
}

type UserServiceInterface interface {
	CreateUser(user CreateUserDTO) uuid.UUID
	//UpdateTask(ix uuid.UUID, newValues Task) (task *Task, err error)
	//DeleteTask(i uuid.UUID) (id uuid.UUID, err error)
	//GetTaskById(id uuid.UUID) (t *Task, err error)
	//GetSortedTasks() []Task
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo}
}

func (service *UserService) CreateUser(user CreateUserDTO) uuid.UUID {

	newUser := User{
		Username: user.Username,
		Password: user.Password,
	}
	id := service.userRepository.SaveUser(newUser)

	return id
}

//
//func (service *TaskService) UpdateTask(id uuid.UUID, newValues Task) (task *Task, err error) {
//	_, err = service.userRepository.UpdateTask(id, newValues)
//
//	task, _, _ = service.userRepository.FindTaskById(id)
//
//	return
//}
//
//func (service *TaskService) DeleteTask(i uuid.UUID) (id uuid.UUID, err error) {
//
//	task, _, err := service.userRepository.FindTaskById(i)
//
//	if err != nil {
//		id1, _ := uuid.NewUUID()
//		return id1, err
//	}
//
//	err = service.userRepository.DeleteTask(task)
//	return
//}
//
//func (service *TaskService) GetTaskById(id uuid.UUID) (t *Task, err error) {
//	t, _, err = service.userRepository.FindTaskById(id)
//
//	return
//}
//
//func (service *TaskService) GetSortedTasks() (tasks []Task) {
//	tasks, err := service.userRepository.FindAllTasks()
//
//	if err != nil {
//		log.Println(err.Error())
//		return make([]Task, 0)
//	}
//
//	return
//}
