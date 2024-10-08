package service

import (
	"github.com/google/uuid"
	. "github.com/vitorcsbrito/go-academy-todo/model"
	. "github.com/vitorcsbrito/go-academy-todo/model/task"
	. "github.com/vitorcsbrito/go-academy-todo/repository/task"
	"github.com/vitorcsbrito/mapper"
	"log"
)

type TaskService struct {
	taskRepository Repository
	userService    UserService
}

type TaskServiceInterface interface {
	CreateTask(description string) (*Task, error)
	UpdateTask(ix uuid.UUID, newValues Task) (task *Task, err error)
	DeleteTask(i uuid.UUID) (id uuid.UUID, err error)
	GetTaskById(id uuid.UUID) (t *Task, err error)
	GetSortedTasks() []Task
}

func NewTaskService(repo Repository, userService *UserService) *TaskService {
	return &TaskService{repo, *userService}
}

func (service *TaskService) CreateTask(taskDto CreateTaskDTO) (*Task, error) {

	user, userErr := service.userService.GetUser(taskDto.UserId)
	if userErr != nil {
		return nil, userErr
	}

	te := mapper.NewEntity(taskDto.Description, &user)

	id, err := service.taskRepository.SaveTask(te)
	if err != nil {
		return nil, err
	}

	createdTask, _ := service.GetTaskById(id)

	return createdTask, nil
}

func (service *TaskService) UpdateTask(id uuid.UUID, newValues Task) (task *Task, err error) {
	_, err = service.taskRepository.UpdateTask(id, newValues)

	task, _, _ = service.taskRepository.FindById(id)

	return
}

func (service *TaskService) DeleteTask(i uuid.UUID) (id uuid.UUID, err error) {

	task, _, err := service.taskRepository.FindById(i)

	if err != nil {
		id1, _ := uuid.NewUUID()
		return id1, err
	}

	err = service.taskRepository.DeleteTask(task)
	return
}

func (service *TaskService) GetTaskById(id uuid.UUID) (t *Task, err error) {
	t, _, err = service.taskRepository.FindById(id)

	return
}

func (service *TaskService) GetSortedTasks() (tasks []Task) {
	tasks, err := service.taskRepository.FindAllTasks()

	if err != nil {
		log.Println(err.Error())
		return make([]Task, 0)
	}

	return
}
