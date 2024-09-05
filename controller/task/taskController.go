package task

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	. "github.com/vitorcsbrito/go-academy-todo/model"
	. "github.com/vitorcsbrito/go-academy-todo/model/task"
	. "github.com/vitorcsbrito/go-academy-todo/service"
	. "github.com/vitorcsbrito/middleware"
	. "github.com/vitorcsbrito/utils/errors"
	. "github.com/vitorcsbrito/utils/requests"
	"gorm.io/gorm"
	"log"
	. "net/http"
)

type Controller struct {
	taskService *TaskService
}

func NewTaskController(taskService *TaskService) *Controller {
	t := &Controller{
		taskService,
	}
	return t
}

func (taskController *Controller) RegisterHandlers(mux *ServeMux) {
	mux.Handle("GET /tasks", Auth(getAllTasks(taskController)))
	mux.Handle("GET /tasks/{id}", Auth(getTaskById(taskController)))
	mux.Handle("POST /tasks", Auth(createTask(taskController)))
	mux.Handle("PUT /tasks/{id}", Auth(updateTask(taskController)))
	mux.Handle("DELETE /tasks/{id}", Auth(deleteTask(taskController)))
}

func getAllTasks(controller *Controller) func(w ResponseWriter, r *Request) {
	return func(w ResponseWriter, r *Request) {
		task := controller.taskService.GetSortedTasks()

		res := struct {
			Tasks []Task `json:"tasks"`
		}{
			Tasks: task,
		}

		NewOkResponse(w, res)
	}
}

func getTaskById(controller *Controller) func(w ResponseWriter, r *Request) {
	return func(w ResponseWriter, r *Request) {
		value := r.URL.Path[len("/tasks/"):]

		if value == "" {
			NewBadRequestResponse(w, ErrTaskIdMissingFromRequest)
			return
		}

		log.Println("GET params were:", value)

		uid := uuid.MustParse(value)
		task, err := controller.taskService.GetTaskById(uid)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			NewNotFoundResponse(w, ErrTaskNotFound)
		} else if err != nil {
			NewInternalErrorResponse(w, err)
		} else {
			NewOkResponse(w, task)
		}
	}
}

func createTask(controller *Controller) func(w ResponseWriter, r *Request) {
	return func(w ResponseWriter, r *Request) {
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

func deleteTask(controller *Controller) func(w ResponseWriter, r *Request) {
	return func(w ResponseWriter, r *Request) {

		value := r.PathValue("id")
		log.Println("GET params were:", value)

		uid := uuid.MustParse(value)
		taskId, err := controller.taskService.DeleteTask(uid)

		if err != nil {
			NewNotFoundResponse(w, err)
		} else {
			NewOkResponse(w, taskId)
		}

	}
}

func updateTask(controller *Controller) func(w ResponseWriter, r *Request) {
	return func(w ResponseWriter, r *Request) {
		value := r.PathValue("id")
		log.Println("GET params were:", value)

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
