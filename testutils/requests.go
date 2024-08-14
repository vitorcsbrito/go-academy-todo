package testutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func NewCreateTaskRequest(body model.Task) *http.Request {
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(body)

	request, _ := http.NewRequest(http.MethodGet, "/tasks", &buf)
	return request
}

func NewGetTaskRequest(id int) (req *http.Request) {
	req, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/tasks/%d", id), nil)
	return
}
