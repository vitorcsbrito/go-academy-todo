package repository

import (
	"fmt"
	. "go-todo-app/model"
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_Remove_LastElement(t *testing.T) {
	filename := getTestDbFileName("Test_Remove_LastElement")
	testRepo := GetInstance(filename)

	testRepo.Save(Task{Id: 0, Description: "do dishes"})
	testRepo.Save(Task{Id: 1, Description: "do laundry"})

	err := testRepo.Delete(1)

	if err != nil {
		t.Helper()
		t.Fatalf("didnt expect an err, but got one")
	}

	_, _, err1 := testRepo.FindById(1)

	if err1 == nil {
		t.Errorf("Expected task to not be found")
		return
	}

	cleanup(filename)
}

func Test_Remove_FirstElement(t *testing.T) {
	filename := getTestDbFileName("Test_Remove_FirstElement")
	testRepo := GetInstance(filename)

	testRepo.Save(Task{Id: 0, Description: "do dishes"})
	testRepo.Save(Task{Id: 1, Description: "do laundry"})

	err := testRepo.Delete(0)

	if err != nil {
		t.Helper()
		t.Fatalf("didnt expect an err, but got one")
	}

	_, _, err1 := testRepo.FindById(0)

	if err1 == nil {
		t.Errorf("Expected task to not be found")
		return
	}

	//cleanup(filename)
}

func Test_SaveTask(t *testing.T) {
	filename := getTestDbFileName("Test_SaveTask")
	testRepo := GetInstance(filename)

	allTasks := testRepo.FindAll()

	prevLen := len(allTasks)

	index := testRepo.Save(Task{Id: 0, Description: "do dishes"})

	allTasks = testRepo.FindAll()
	postLen := len(allTasks)

	if postLen <= prevLen {
		t.Helper()
		t.Fatalf("task didnt save")
	}

	if index != 0 {
		t.Helper()
		t.Fatalf("didnt expect an err, but got one")
	}

	task, _, err1 := testRepo.FindById(0)

	if (*task).Description != "do dishes" {
		t.Helper()
		t.Fatalf("task has wrong description")
	}

	if err1 != nil {
		t.Errorf("Expected task to not be found")
		return
	}

	cleanup(filename)
}

func TestFindById(t *testing.T) {
	filename := getTestDbFileName("TestFindById")
	testRepo := GetInstance(filename)

	testRepo.Save(Task{Id: 0, Description: "do dishes"})
	testRepo.Save(Task{Id: 1, Description: "do laundry"})

	type args struct {
		index int
	}

	tests := []struct {
		name      string
		args      args
		wantTask  Task
		wantIndex int
		wantErr   bool
	}{
		{name: "Finds existing task", args: args{0}, wantTask: Task{Id: 0, Description: "do dishes"}, wantIndex: 0, wantErr: false},
		{name: "Finds existing task", args: args{1}, wantTask: Task{Id: 1, Description: "do laundry"}, wantIndex: 1, wantErr: false},
		{name: "Throws error with unknown task", args: args{2}, wantTask: Task{Id: -1, Description: ""}, wantIndex: -1, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTask, gotIndex, err := testRepo.FindById(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(*gotTask, tt.wantTask) {
				t.Errorf("FindById() gotTask = %v, wantTask %v", gotTask, tt.wantTask)
			}
			if gotIndex != tt.wantIndex {
				t.Errorf("FindById() gotIndex = %v, wantTask %v", gotIndex, tt.wantIndex)
			}
		})
	}

	cleanup(filename)
}

func getTestDbFileName(testName string) string {
	return fmt.Sprintf("test_%s_%d.json", testName, time.Now().Unix())
}

func cleanup(filename string) {
	err := os.Remove(filename)

	if err != nil {
		panic(err)
	}
}
