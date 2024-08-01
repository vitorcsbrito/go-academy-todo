package service

import (
	"errors"
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

	// Call BindJSON to bind the received JSON to
	// newTask.
	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	newTask.Id = findLatestId()

	Tasks = append(Tasks, newTask)
	files.WriteTasksToJsonFile("tasks.json", Tasks...)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func findTask(id int) (Task, int, error) {

	for i, todo := range Tasks {
		if todo.Id == id {
			return todo, i, nil
		}
	}
	return Task{-1, "", false}, -1, errors.New("math: square root of negative number")
}

func UpdateTask(c *gin.Context) {
	var newValues Task
	if err := c.BindJSON(&newValues); err != nil {
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	_, i, _ := findTask(id)

	Tasks[i].Description = newValues.Description
	Tasks[i].Done = newValues.Done

	//tasks = append(tasks, newValues)
	files.WriteTasksToJsonFile("tasks.json", Tasks...)
	c.IndentedJSON(http.StatusCreated, Tasks[i])
}

func DeleteTask(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	_, i, _ := findTask(id)

	Tasks = remove(Tasks, i)

	//Tasks = append(Tasks, newValues)
	files.WriteTasksToJsonFile("Tasks.json", Tasks...)
	c.IndentedJSON(http.StatusCreated, Tasks[i])
}

func remove(slice []Task, s int) []Task {
	lastIndex := len(slice) - 1

	if len(slice) > 0 && s == lastIndex {
		slice = slice[:len(slice)-1]
		return slice
	} else {
		return append(slice[:s], slice[s+1:]...)
	}
}

func findLatestId() int {
	if len(Tasks) == 0 {
		return 0
	}

	sort.SliceStable(Tasks, func(i, j int) bool {
		return Tasks[i].Id > Tasks[j].Id
	})

	return Tasks[0].Id + 1
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
