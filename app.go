package main

import (
	. "go-todo-app/controller"
	"go-todo-app/repository"
	"go-todo-app/service"
	"html/template"
	"log"
	"net/http"
)

func main() {
	taskRepository := repository.GetInstance()
	taskService := service.NewTaskService(taskRepository)
	userService := service.NewUserService(taskRepository)

	tmpl := template.Must(template.ParseFiles("templates/tasks.html"))

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, taskService.GetSortedTasks())
	})

	http.HandleFunc("GET /tasks/{id}", GetTaskById(taskService))
	http.HandleFunc("POST /tasks", CreateTask(taskService))
	http.HandleFunc("PUT /tasks/{id}", UpdateTask(taskService))
	http.HandleFunc("DELETE /tasks/{id}", DeleteTask(taskService))

	http.HandleFunc("POST /users", CreateUser(userService))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
