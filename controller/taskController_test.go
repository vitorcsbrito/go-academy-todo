package controller

import (
	. "go-todo-app/model"
	. "go-todo-app/repository"
	. "go-todo-app/service"
	. "go-todo-app/testutils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateTask(t *testing.T) {
	ts, fn := SetupTaskService()

	t.Run("get task by id 1", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/tasks/1", nil)
		response := httptest.NewRecorder()

		GetTaskById(ts)(response, request)

		got := response.Body.String()
		want := "{\"id\":0,\"description\":\"do dishes\",\"done\":false}\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	Cleanup(fn)
}
