package main

import (
	"log"
	"net/http"

	taskRepo "github.com/vitorcsbrito/go-academy-todo/repository/task/mysql"
	userRepo "github.com/vitorcsbrito/go-academy-todo/repository/user/mysql"

	. "github.com/vitorcsbrito/go-academy-todo/controller"
	"github.com/vitorcsbrito/go-academy-todo/repository"
	. "github.com/vitorcsbrito/go-academy-todo/service"
)

func main() {
	database := repository.GetInstance()

	userRepository := userRepo.NewMySqlRepository(database)
	userService := NewUserService(userRepository)
	userController := NewUserController(userService)

	taskRepository := taskRepo.NewMySqlRepository(database)
	taskService := NewTaskService(taskRepository, userService)
	taskController := NewTaskController(taskService)

	// Tasks
	http.HandleFunc("GET /", RenderInterface(taskController))
	http.HandleFunc("GET /tasks", GetAllTasks(taskController))
	http.HandleFunc("GET /tasks/{id}", GetTaskById(taskController))
	http.HandleFunc("POST /tasks", CreateTask(taskController))
	http.HandleFunc("PUT /tasks/{id}", UpdateTask(taskController))
	http.HandleFunc("DELETE /tasks/{id}", DeleteTask(taskController))

	// Users
	http.HandleFunc("POST /users", CreateUser(userController))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
