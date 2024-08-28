package service

import (
	"github.com/google/uuid"
	model2 "github.com/vitorcsbrito/go-academy-todo/model"
	model "github.com/vitorcsbrito/go-academy-todo/model/user"
	"github.com/vitorcsbrito/go-academy-todo/repository/user"
	"github.com/vitorcsbrito/mapper"
)

type UserService struct {
	userRepository user.Repository
}

type UserServiceInterface interface {
	CreateUser(user model.CreateUserDTO) (uuid.UUID, error)
	GetUser(id uuid.UUID) (model2.User, error)
	//UpdateTask(ix uuid.UUID, newValues Task) (task *Task, err error)
	//DeleteTask(i uuid.UUID) (id uuid.UUID, err error)
	//GetTaskById(id uuid.UUID) (t *Task, err error)
	//GetSortedTasks() []task.Task
}

func NewUserService(repo user.Repository) *UserService {
	return &UserService{repo}
}

func (service *UserService) CreateUser(userDto model.CreateUserDTO) (uuid.UUID, error) {

	newUser := mapper.DtoToEntityNewUser(userDto)

	id, err := service.userRepository.Save(newUser)

	return id, err
}

func (service *UserService) GetUser(id uuid.UUID) (model2.User, error) {
	user, err := service.userRepository.Get(id)

	return user, err
}

//
//func (service *TaskService) UpdateTask(id uuid.UUID, newValues Task) (task *Task, err error) {
//	_, err = service.userRepository.UpdateTask(id, newValues)
//
//	task, _, _ = service.userRepository.FindById(id)
//
//	return
//}
//
//func (service *TaskService) DeleteTask(i uuid.UUID) (id uuid.UUID, err error) {
//
//	task, _, err := service.userRepository.FindById(i)
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
//	t, _, err = service.userRepository.FindById(id)
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
