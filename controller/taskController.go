package controller

import (
	"encoding/json"
	"fmt"
	. "go-todo-app/errors"
	. "go-todo-app/model"
	"go-todo-app/service"
	"net/http"
)

func GetTaskById(taskService *service.TaskService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		value := r.URL.Path[len("/tasks/"):]

		if value == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{
				Message: "task id is mandatory",
			})
			return
		}

		fmt.Println("GET params were:", value)

		task, err := taskService.GetTaskById(value)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)

			json.NewEncoder(w).Encode(ErrorResponse{
				Message: err.Error(),
			})
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(task)
		}
	}
}

func CreateTask(taskService *service.TaskService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		if body == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			json.NewEncoder(w).Encode(ErrorResponse{
				Message: "Missing task description",
			})
			return
		}

		var task Task
		err := json.NewDecoder(body).Decode(&task)

		newTask := taskService.CreateTask(task)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)

			json.NewEncoder(w).Encode(ErrorResponse{
				Message: err.Error(),
			})
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(newTask)
		}
	}
}

func DeleteTask(taskService *service.TaskService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		value := r.PathValue("id")
		fmt.Println("GET params were:", value)

		taskId, err := taskService.DeleteTask(value)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)

			json.NewEncoder(w).Encode(ErrorResponse{
				Message: err.Error(),
			})
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(taskId)
		}

	}
}

func UpdateTask(taskService *service.TaskService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		value := r.PathValue("id")
		fmt.Println("GET params were:", value)

		decoder := json.NewDecoder(r.Body)

		var newTask Task
		err1 := decoder.Decode(&newTask)

		if err1 != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			json.NewEncoder(w).Encode(ErrorResponse{
				Message: err1.Error(),
			})
		}

		updatedTask, _ := taskService.UpdateTask(value, newTask)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(updatedTask)

	}
}
