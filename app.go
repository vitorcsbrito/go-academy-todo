package main

import (
	"github.com/gin-gonic/gin"
	"go-todo-app/repository"
	"go-todo-app/service"
	"net/http"
)

func main() {
	taskRepository := repository.GetInstance("tasks.json")
	taskService := service.NewTaskService(taskRepository)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tasks.tmpl", taskService.GetSortedTasks())
	})

	router.GET("/tasks/:id", taskService.GetTaskById)
	router.POST("/tasks", taskService.CreateTask)
	router.PUT("/tasks/:id", taskService.UpdateTask)
	router.DELETE("/tasks/:id", taskService.DeleteTask)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
