package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	. "github.com/vitorcsbrito/go-academy-todo/model"
	. "github.com/vitorcsbrito/go-academy-todo/model/task"
	. "github.com/vitorcsbrito/go-academy-todo/service"
	. "github.com/vitorcsbrito/utils/errors"
	. "github.com/vitorcsbrito/utils/requests"
	"gorm.io/gorm"
	"html/template"
	"net/http"
)

type TaskController struct {
	taskService *TaskService
}

func NewTaskController(taskService *TaskService) *TaskController {
	return &TaskController{taskService}
}

func RenderInterface(controller *TaskController) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/tasks.html"))
		tmpl.Execute(w, controller.taskService.GetSortedTasks())
	}
}

func GetAllTasks(controller *TaskController) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		task := controller.taskService.GetSortedTasks()

		NewOkResponse(w, task)
	}
}

func GetTaskById(controller *TaskController) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		value := r.URL.Path[len("/tasks/"):]

		if value == "" {
			NewBadRequestResponse(w, ErrTaskIdMissingFromRequest)
			return
		}

		fmt.Println("GET params were:", value)

		uid := uuid.MustParse(value)
		task, err := controller.taskService.GetTaskById(uid)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			NewNotFoundResponse(w, err)
		} else if err != nil {
			NewInternalErrorResponse(w, err)
		} else {
			NewOkResponse(w, task)
		}
	}
}

func CreateTask(controller *TaskController) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		if body == nil {
			NewNotFoundResponse(w, ErrTaskDescriptionNotFound)
			return
		}

		var task CreateTaskDTO
		err := json.NewDecoder(body).Decode(&task)

		newTask, createErr := controller.taskService.CreateTask(task)

		if err != nil {
			NewNotFoundResponse(w, err)
		} else if createErr != nil {
			NewBadRequestResponse(w, createErr)
		} else {
			NewOkResponse(w, newTask)
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
			NewNotFoundResponse(w, err)
		} else {
			NewOkResponse(w, taskId)
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
			NewBadRequestResponse(w, err1)
		}

		uid := uuid.MustParse(value)
		updatedTask, _ := controller.taskService.UpdateTask(uid, newTask)

		NewOkResponse(w, updatedTask)
	}
}
