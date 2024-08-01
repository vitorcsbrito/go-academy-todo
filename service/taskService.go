package service

import (
	"github.com/gin-gonic/gin"
	"go-todo-app/files"
	. "go-todo-app/model"
	. "go-todo-app/repository"
	"net/http"
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

	Tasks = append(Tasks, newTask)
	files.WriteTasksToJsonFile("tasks.json", Tasks...)
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

	Tasks[i].Description = newValues.Description
	Tasks[i].Done = newValues.Done

	files.WriteTasksToJsonFile("tasks.json", Tasks...)
	c.IndentedJSON(http.StatusCreated, Tasks[i])
}

func DeleteTask(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	_, i, err := FindById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	}

	Tasks = Delete(Tasks, i)

	files.WriteTasksToJsonFile("Tasks.json", Tasks...)
	c.IndentedJSON(http.StatusOK, "")
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)

	for _, a := range Tasks {
		if a.Id == i {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
