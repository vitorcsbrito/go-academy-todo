package task

import (
	"github.com/vitorcsbrito/go-academy-todo/repository"
	tr "github.com/vitorcsbrito/go-academy-todo/repository/task/mysql"
	ur "github.com/vitorcsbrito/go-academy-todo/repository/user/mysql"
	"github.com/vitorcsbrito/go-academy-todo/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAllTasks(t *testing.T) {

	db := repository.GetInstance()
	ur := ur.NewMySqlRepository(db)
	tr := tr.NewMySqlRepository(db)

	userService := service.NewUserService(ur)
	taskService := service.NewTaskService(tr, userService)
	taskController := NewTaskController(taskService)

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAllTasks(taskController))

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"tasks":[]}`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}
