package testutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/vitorcsbrito/go-academy-todo/model/task"
	"net/http"
)

func NewCreateTaskRequest(body task.Task) *http.Request {
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(body)

	request, _ := http.NewRequest(http.MethodGet, "/tasks", &buf)
	return request
}

func NewGetTaskRequest(id int) (req *http.Request) {
	req, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/tasks/%d", id), nil)
	return
}
