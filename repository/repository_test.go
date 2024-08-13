package repository

import (
	. "go-todo-app/model"
	. "go-todo-app/testutils"
	"testing"
)

func Test_Remove_LastElement(t *testing.T) {
	testRepo, filename := SetupTaskRepository()
	err := testRepo.Delete(1)

	if err != nil {
		t.Helper()
		t.Errorf("didnt expect an err, but got one")
	}

	_, _, err1 := testRepo.FindById(1)

	if err1 == nil {
		t.Errorf("Expected task to not be found")
		return
	}

	Cleanup(filename)
}

func Test_Remove_FirstElement(t *testing.T) {
	testRepo, filename := SetupTaskRepository()

	err := testRepo.Delete(0)

	if err != nil {
		t.Helper()
		t.Errorf("didnt expect an err, but got one")
	}

	_, _, err1 := testRepo.FindById(0)

	if err1 == nil {
		t.Errorf("Expected task to not be found")
		return
	}

	Cleanup(filename)
}

func Test_SaveTask(t *testing.T) {
	testRepo, filename := SetupTaskRepository()

	allTasks := testRepo.FindAll()

	prevLen := len(allTasks)

	index := testRepo.Save(Task{Description: "do dishes"})

	allTasks = testRepo.FindAll()
	postLen := len(allTasks)

	if postLen <= prevLen {
		t.Helper()
		t.Errorf("task didnt save")
	}

	task, _, err1 := testRepo.FindById(index)

	if (*task).Description != "do dishes" {
		t.Helper()
		t.Errorf("task has wrong description")
	}

	if err1 != nil {
		t.Errorf("Expected task to not be found")
		return
	}

	Cleanup(filename)
}

func TestFindById(t *testing.T) {
	testRepo, filename := SetupTaskRepository()

	taskId := testRepo.findLatestId() - 1

	gotTask, gotIndex, err := testRepo.FindById(taskId)
	if err != nil {
		t.Errorf("FindById() error = %v, wantErr %v", err, "tt.wantErr")
		return
	}

	if taskId != gotTask.Id {
		t.Errorf("FindById() gotIndex = %v, wantIndex %v", gotIndex, taskId)
	}

	Cleanup(filename)
}

func SetupTaskRepository() (*Repository, string) {
	filename := GetTestDbFileName()
	testRepo := GetInstance()

	testRepo.Save(Task{Id: 0, Description: "do dishes"})
	testRepo.Save(Task{Id: 1, Description: "do laundry"})

	return testRepo, filename
}
