package controller

import (
	"encoding/json"
	. "go-todo-app/errors"
	. "go-todo-app/model"
	"go-todo-app/service"
	"net/http"
)

func CreateUser(userService *service.UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		if body == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			json.NewEncoder(w).Encode(ErrorResponse{
				Message: "Missing username & password",
			})
			return
		}

		var user CreateUserDTO
		err := json.NewDecoder(body).Decode(&user)

		newUser := userService.CreateUser(user)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)

			json.NewEncoder(w).Encode(ErrorResponse{
				Message: err.Error(),
			})
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(newUser)
		}
	}
}

//func GetTaskById(taskService *service.TaskService) func(w http.ResponseWriter, r *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		value := r.URL.Path[len("/tasks/"):]
//
//		if value == "" {
//			w.WriteHeader(http.StatusBadRequest)
//			json.NewEncoder(w).Encode(ErrorResponse{
//				Message: "task id is mandatory",
//			})
//			return
//		}
//
//		fmt.Println("GET params were:", value)
//
//		uid := uuid.MustParse(value)
//		task, err := taskService.GetTaskById(uid)
//
//		if err != nil {
//			w.Header().Set("Content-Type", "application/json")
//			w.WriteHeader(http.StatusNotFound)
//
//			json.NewEncoder(w).Encode(ErrorResponse{
//				Message: err.Error(),
//			})
//		} else {
//			w.Header().Set("Content-Type", "application/json")
//			w.WriteHeader(http.StatusOK)
//
//			json.NewEncoder(w).Encode(task)
//		}
//	}
//}
//
//
//func DeleteTask(taskService *service.TaskService) func(w http.ResponseWriter, r *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//
//		value := r.PathValue("id")
//		fmt.Println("GET params were:", value)
//
//		uid := uuid.MustParse(value)
//		taskId, err := taskService.DeleteTask(uid)
//
//		if err != nil {
//			w.Header().Set("Content-Type", "application/json")
//			w.WriteHeader(http.StatusNotFound)
//
//			json.NewEncoder(w).Encode(ErrorResponse{
//				Message: err.Error(),
//			})
//		} else {
//			w.Header().Set("Content-Type", "application/json")
//			w.WriteHeader(http.StatusOK)
//
//			json.NewEncoder(w).Encode(taskId)
//		}
//
//	}
//}
//
//func UpdateTask(taskService *service.TaskService) func(w http.ResponseWriter, r *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		value := r.PathValue("id")
//		fmt.Println("GET params were:", value)
//
//		decoder := json.NewDecoder(r.Body)
//
//		var newTask Task
//		err1 := decoder.Decode(&newTask)
//
//		if err1 != nil {
//			w.Header().Set("Content-Type", "application/json")
//			w.WriteHeader(http.StatusBadRequest)
//
//			json.NewEncoder(w).Encode(ErrorResponse{
//				Message: err1.Error(),
//			})
//		}
//
//		uid := uuid.MustParse(value)
//		updatedTask, _ := taskService.UpdateTask(uid, newTask)
//
//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusOK)
//
//		json.NewEncoder(w).Encode(updatedTask)
//
//	}
//}
