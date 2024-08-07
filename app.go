package main

import (
	. "go-todo-app/controller"
	"go-todo-app/repository"
	"go-todo-app/service"
	"html/template"
	"net/http"
)

func main() {
	taskRepository := repository.GetInstance("tasks.json")
	taskService := service.NewTaskService(taskRepository)

	tmpl := template.Must(template.ParseFiles("templates/tasks.tmpl"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := taskService.GetSortedTasks()
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/tasks/{id}", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			GetTaskById(taskService)(writer, request)
		case http.MethodDelete:
			DeleteTask(taskService)(writer, request)
		case http.MethodPut:
			UpdateTask(taskService)(writer, request)
		}
	})

	//handler := http.HandlerFunc(GetTaskById(taskService))

	http.HandleFunc("/tasks", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			CreateTask(taskService)(writer, request)
		}
	})

	http.ListenAndServe(":8080", nil)
}
