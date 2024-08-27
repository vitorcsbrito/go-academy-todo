package controller

import (
	"encoding/json"
	"errors"
	"github.com/go-sql-driver/mysql"
	. "github.com/vitorcsbrito/go-academy-todo/model/user"
	"github.com/vitorcsbrito/go-academy-todo/service"
	. "github.com/vitorcsbrito/utils/errors"
	. "github.com/vitorcsbrito/utils/requests"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService,
	}
}

func CreateUser(userController *UserController) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		if body == nil {
			NewBadRequestResponse(w, ErrMissingErrorDetails)
			return
		}

		var user CreateUserDTO
		err := json.NewDecoder(body).Decode(&user)

		newUser, createUserErr := userController.userService.CreateUser(user)

		var mySqlError *mysql.MySQLError
		if err != nil {
			NewBadRequestResponse(w, err)
		} else if errors.As(createUserErr, &mySqlError) {
			NewBadRequestResponse(w, ErrEmailTaken)
		} else {
			NewOkResponse(w, newUser)
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
