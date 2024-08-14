package controller

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	. "model"
	"net/http"
	"service"
	"todoerrors"
)

type TaskController struct {
	taskService *service.TaskService
}

func NewTaskController(taskService *service.TaskService) *TaskController {
	return &TaskController{taskService}
}

func GetTaskById(controller *TaskController) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		value := r.URL.Path[len("/tasks/"):]

		if value == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(todoerrors.NewErrTaskNotFound(value))
			return
		}

		fmt.Println("GET params were:", value)

		uid := uuid.MustParse(value)
		task, err := controller.taskService.GetTaskById(uid)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)

			json.NewEncoder(w).Encode(todoerrors.NewErrResponse(value))
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(task)
		}
	}
}

func CreateTask(controller *TaskController) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		if body == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			json.NewEncoder(w).Encode(todoerrors.NewErrResponse("Missing task description"))
			return
		}

		var task CreateTaskDTO
		err := json.NewDecoder(body).Decode(&task)

		newTask := controller.taskService.CreateTask(task.Description)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)

			json.NewEncoder(w).Encode(todoerrors.NewErrResponse(err.Error()))
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(newTask)
		}
	}
}

func DeleteTask(controller *TaskController) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		value := r.PathValue("id")
		fmt.Println("GET params were:", value)

		uid := uuid.MustParse(value)
		taskId, err := controller.taskService.DeleteTask(uid)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)

			json.NewEncoder(w).Encode(todoerrors.NewErrResponse(value))
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(taskId)
		}

	}
}

func UpdateTask(controller *TaskController) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		value := r.PathValue("id")
		fmt.Println("GET params were:", value)

		decoder := json.NewDecoder(r.Body)

		var newTask Task
		err1 := decoder.Decode(&newTask)

		if err1 != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			json.NewEncoder(w).Encode(todoerrors.NewErrResponse(err1.Error()))
		}

		uid := uuid.MustParse(value)
		updatedTask, _ := controller.taskService.UpdateTask(uid, newTask)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(updatedTask)

	}
}
