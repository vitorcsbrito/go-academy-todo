package service

import (
	"github.com/gin-gonic/gin"
	. "go-todo-app/model"
	. "go-todo-app/repository"
	"net/http"
	"sort"
	"strconv"
)

type TaskService struct {
	repo InterfaceRepository
}

func NewTaskService(repo InterfaceRepository) *TaskService {
	return &TaskService{repo}
}

func (service *TaskService) CreateTask(c *gin.Context) {
	var newTask Task

	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	service.repo.Save(newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func (service *TaskService) UpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var newValues Task
	if err := c.BindJSON(&newValues); err != nil {
		return
	}

	i, err := service.repo.Update(id, newValues)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	}

	task, _, _ := service.repo.FindById(i)
	c.IndentedJSON(http.StatusCreated, task)
}

func (service *TaskService) DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := service.repo.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	}

	c.IndentedJSON(http.StatusOK, "")
}

func (service *TaskService) GetTaskById(c *gin.Context) {
	i, _ := strconv.Atoi(c.Param("id"))

	t, _, err := service.repo.FindById(i)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	}

	c.IndentedJSON(http.StatusOK, t)

}

func (service *TaskService) GetSortedTasks() []Task {
	tasks := service.repo.FindAll()
	sort.SliceStable(tasks, func(i, j int) bool {
		return tasks[i].Id < tasks[j].Id
	})

	return tasks
}
