package controller

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	smt "go-todo-app/entities"
	. "go-todo-app/model"
	. "go-todo-app/repository"
	. "go-todo-app/service"
	. "go-todo-app/testutils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTaskById(t *testing.T) {
	ts, fn := SetupTaskService()

	t.Run("get task by id 1", func(t *testing.T) {
		request := NewGetTaskRequest(1)
		response := httptest.NewRecorder()

		GetTaskById(ts)(response, request)

		got := response.Body.String()
		want := "{\"id\":1,\"description\":\"do laundry\",\"done\":false}\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("get task by id for unknown task", func(t *testing.T) {
		request := NewGetTaskRequest(4)
		response := httptest.NewRecorder()

		GetTaskById(ts)(response, request)

		got := response.Body.String()
		want := "{\"message\":\"no task found for id 4\"}\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("get task by id with missing id", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/tasks/"), nil)
		response := httptest.NewRecorder()

		GetTaskById(ts)(response, request)

		got := response.Body.String()
		want := "{\"message\":\"task id is mandatory\"}\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	Cleanup(fn)
}

func TestCreateTask(t *testing.T) {
	mockRepo, fn := SetupMockRepository()
	ts := NewTaskService(mockRepo)

	t.Run("bad request with missing task description", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
		response := httptest.NewRecorder()

		CreateTask(ts)(response, request)

		got := response.Body.String()
		want := "{\"message\":\"Missing task description\"}\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("happy path", func(t *testing.T) {
		body := Task{Description: "task"}

		request := NewCreateTaskRequest(body)
		response := httptest.NewRecorder()

		CreateTask(ts)(response, request)

		var gotTask Task
		_ = json.Unmarshal([]byte(response.Body.String()), &gotTask)

		if gotTask.Description != body.Description {
			t.Errorf("got %q, want %q", gotTask.Description, body.Description)
		}
	})

	Cleanup(fn)
}

func SetupTaskService() (taskService *TaskService, fn string) {
	testRepo, fn := SetupTaskRepository()

	taskService = NewTaskService(testRepo)
	return
}

func SetupTaskRepository() (*Repository, string) {
	filename := GetTestDbFileName()
	testRepo := GetInstance()

	id1, _ := uuid.NewUUID()
	id2, _ := uuid.NewUUID()

	testRepo.SaveTask(smt.TaskEntity{Id: id1, Description: "do dishes"})
	testRepo.SaveTask(smt.TaskEntity{Id: id2, Description: "do laundry"})

	return testRepo, filename
}

func SetupMockRepository() (*Repository, string) {
	//taskArr := make([]Task, 0)
	filename := GetTestDbFileName()
	mockRepo := &Repository{
		//Tasks:    &taskArr,
		//
		//Filename: filename,
	}
	return mockRepo, filename
}
