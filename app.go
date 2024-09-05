package main

import (
	"github.com/vitorcsbrito/go-academy-todo/controller/task"
	"github.com/vitorcsbrito/go-academy-todo/controller/user"
	"github.com/vitorcsbrito/go-academy-todo/repository"
	taskRepo "github.com/vitorcsbrito/go-academy-todo/repository/task/mysql"
	userRepo "github.com/vitorcsbrito/go-academy-todo/repository/user/mysql"
	. "github.com/vitorcsbrito/go-academy-todo/service"
	"github.com/vitorcsbrito/middleware"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	wrappedMux := middleware.RequestLogger(mux)

	database := repository.GetInstance()

	userRepository := userRepo.NewMySqlRepository(database)
	taskRepository := taskRepo.NewMySqlRepository(database)

	userService := NewUserService(userRepository)
	taskService := NewTaskService(taskRepository, userService)

	userController := user.NewUserController(userService)
	taskController := task.NewTaskController(taskService)

	userController.RegisterHandlers(mux)
	taskController.RegisterHandlers(mux)

	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}
