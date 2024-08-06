package service

import (
	. "go-todo-app/model"
	. "go-todo-app/repository"
	"sort"
	"strconv"
)

type TaskService struct {
	repo InterfaceRepository
}

func NewTaskService(repo InterfaceRepository) *TaskService {
	return &TaskService{repo}
}

func (service *TaskService) CreateTask(task Task) {
	service.repo.Save(task)
}

func (service *TaskService) UpdateTask(ix string, newValues Task) (task *Task, err error) {
	id, _ := strconv.Atoi(ix)

	_, err = service.repo.Update(id, newValues)

	task, _, _ = service.repo.FindById(id)
	return
}

func (service *TaskService) DeleteTask(i string) (id int, err error) {
	id, _ = strconv.Atoi(i)

	err = service.repo.Delete(id)
	return
}

func (service *TaskService) GetTaskById(id string) (t *Task, err error) {
	i, _ := strconv.Atoi(id)
	t, _, err = service.repo.FindById(i)

	return
}

func (service *TaskService) GetSortedTasks() []Task {
	tasks := service.repo.FindAll()

	sort.SliceStable(tasks, func(i, j int) bool {
		return (tasks)[i].Id < (tasks)[j].Id
	})

	return tasks
}
