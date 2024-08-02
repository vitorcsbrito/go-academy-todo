package main

import (
	"github.com/gin-gonic/gin"
	"go-todo-app/files"
	. "go-todo-app/repository"
	"go-todo-app/service"
	"net/http"
)

func main() {
	filename := "tasks.json"

	files.ReadTasksFromJson(filename)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tasks.tmpl", GetInstance().Tasks)
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
