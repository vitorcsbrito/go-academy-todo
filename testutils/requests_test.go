package testutils

import (
	"github.com/google/uuid"
	"github.com/vitorcsbrito/go-academy-todo/model"
	"regexp"
	"testing"
	"time"
)

func TestNewCreateTaskRequest(t *testing.T) {

	now := time.Now()
	taskId, _ := uuid.NewUUID()
	userId, _ := uuid.NewUUID()
	taskModel := model.Task{
		ID:          taskId,
		Description: "smt",
		Done:        false,
		UserId:      userId,
		CreatedAt:   &now,
	}

	request := NewCreateTaskRequest(taskModel)

	wantPath := regexp.MustCompile("/tasks")
	path := request.URL.Path

	if !wantPath.MatchString(path) {
		t.Fatalf(`Expected path to be [%s] but got [%v]`, wantPath, path)
	}

	wantMethod := regexp.MustCompile("GET")
	method := request.Method

	if !wantMethod.MatchString(method) {
		t.Fatalf(`Expected method to be [%s] but got [%v]`, wantMethod, method)
	}
}
