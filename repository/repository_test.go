package repository

import (
	. "go-todo-app/model"
	"testing"
)

var tasks []Task

func TestRemoveLastElement(t *testing.T) {
	tasks := append(tasks,
		Task{Id: 0, Description: "do dishes"},
		Task{Id: 1, Description: "do laundry"})

	tmp := Delete(tasks, 1)

	if len(tmp) != 1 {
		t.Helper()
		t.Fatalf("didnt expect an err, but got one")
	}
}
