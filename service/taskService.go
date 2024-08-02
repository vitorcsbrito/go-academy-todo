package service

import (
	"github.com/gin-gonic/gin"
	"go-todo-app/files"
	. "go-todo-app/model"
	. "go-todo-app/repository"
	"net/http"
	"sort"
	"strconv"
)

// CreateTask adds an album from JSON received in the request body.
func CreateTask(c *gin.Context) {
	var newTask Task

	// Call BindJSON to bind the received JSON to newTask.
	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	newTask.Id = FindLatestId()

	GetInstance().Tasks = append(GetInstance().Tasks, newTask)
	files.WriteTasksToJsonFile("tasks.json")
	c.IndentedJSON(http.StatusCreated, newTask)
}

func UpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, i, err := FindById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	}

	var newValues Task
	if err := c.BindJSON(&newValues); err != nil {
		return
	}

	GetInstance().Tasks[i].Description = newValues.Description
	GetInstance().Tasks[i].Done = newValues.Done

	files.WriteTasksToJsonFile("tasks.json")
	c.IndentedJSON(http.StatusCreated, GetInstance().Tasks[i])
}

func DeleteTask(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	_, i, err := FindById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	}

	GetInstance().Tasks = Delete(i)

	files.WriteTasksToJsonFile("Tasks.json")
	c.IndentedJSON(http.StatusOK, "")
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)

	for _, a := range GetInstance().Tasks {
		if a.Id == i {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func GetSortedTasks() []Task {
	sort.SliceStable(GetInstance().Tasks, func(i, j int) bool {
		return GetInstance().Tasks[i].Id < GetInstance().Tasks[j].Id
	})

	return GetInstance().Tasks
}
