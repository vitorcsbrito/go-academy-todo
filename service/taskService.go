package service

import (
	"github.com/google/uuid"
	. "go-todo-app/model"
	. "go-todo-app/repository"
	"log"
)

type TaskService struct {
	repo InterfaceRepository
}

type TaskServiceInterface interface {
	CreateTask(description string) *Task
	UpdateTask(ix uuid.UUID, newValues Task) (task *Task, err error)
	DeleteTask(i uuid.UUID) (id uuid.UUID, err error)
	GetTaskById(id uuid.UUID) (t *Task, err error)
	GetSortedTasks() []Task
}

func NewTaskService(repo InterfaceRepository) *TaskService {
	return &TaskService{repo}
}

func (service *TaskService) CreateTask(description string) *Task {

	te := newEntity(description)
	id := service.repo.Save(te)

	createdTask, _ := service.GetTaskById(id)

	return createdTask
}

func (service *TaskService) UpdateTask(id uuid.UUID, newValues Task) (task *Task, err error) {
	_, err = service.repo.Update(id, newValues)

	task, _, _ = service.repo.FindById(id)

	return
}

func (service *TaskService) DeleteTask(i uuid.UUID) (id uuid.UUID, err error) {

	task, _, err := service.repo.FindById(i)

	if err != nil {
		id1, _ := uuid.NewUUID()
		return id1, err
	}

	err = service.repo.Delete(task)
	return
}

func (service *TaskService) GetTaskById(id uuid.UUID) (t *Task, err error) {
	t, _, err = service.repo.FindById(id)

	return
}

func (service *TaskService) GetSortedTasks() (tasks []Task) {
	tasks, err := service.repo.FindAll()

	if err != nil {
		log.Println(err.Error())
		return make([]Task, 0)
	}

	return
}
