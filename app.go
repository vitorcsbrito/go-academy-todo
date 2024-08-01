package main

import (
	"github.com/gin-gonic/gin"
	"go-todo-app/files"
	. "go-todo-app/model"
	. "go-todo-app/repository"
	"go-todo-app/service"
	"net/http"
	"sort"
)

func main() {
	filename := "tasks.json"

	files.ReadTasksFromJson(filename, &Tasks)

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tasks.tmpl", getSortedTasks())
	})

	router.GET("/tasks/:id", service.GetTaskById)
	router.POST("/tasks", service.CreateTask)
	router.PUT("/tasks/:id", service.UpdateTask)
	router.DELETE("/tasks/:id", service.DeleteTask)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}

func getSortedTasks() []Task {
	sort.SliceStable(Tasks, func(i, j int) bool {
		return Tasks[i].Id < Tasks[j].Id
	})

	return Tasks
}
