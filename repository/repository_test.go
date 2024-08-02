package repository

import (
	. "go-todo-app/model"
	"reflect"
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

func TestFindById(t *testing.T) {
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
		{name: "smt", args: args{0}, wantTask: Task{Id: 0, Description: "123"}, wantIndex: 0, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := FindById(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantTask) {
				t.Errorf("FindById() got = %v, wantTask %v", got, tt.wantTask)
			}
			if got1 != tt.wantIndex {
				t.Errorf("FindById() got1 = %v, wantTask %v", got1, tt.wantIndex)
			}
		})
	}
}
