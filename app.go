package main

import (
	. "controller"
	"html/template"
	"log"
	"net/http"
	. "repository"
	"service"
)

func main() {
	database := GetInstance()
	taskService := service.NewTaskService(database)
	taskController := NewTaskController(taskService)

	userService := service.NewUserService(database)
	userController := NewUserController(userService)

	tmpl := template.Must(template.ParseFiles("templates/tasks.html"))

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, taskService.GetSortedTasks())
	})

	http.HandleFunc("GET /tasks/{id}", GetTaskById(taskController))
	http.HandleFunc("POST /tasks", CreateTask(taskController))
	http.HandleFunc("PUT /tasks/{id}", UpdateTask(taskController))
	http.HandleFunc("DELETE /tasks/{id}", DeleteTask(taskController))

	http.HandleFunc("POST /users", CreateUser(userController))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
