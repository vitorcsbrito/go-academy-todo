package repository

import (
	. "go-todo-app/model"
	"reflect"
	"testing"
)

func TestRemoveLastElement(t *testing.T) {
	GetInstance().tasks = append(GetInstance().tasks,
		Task{Id: 0, Description: "do dishes"},
		Task{Id: 1, Description: "do laundry"})

	tmp := Delete(1)

	if len(tmp) != 1 {
		t.Helper()
		t.Fatalf("didnt expect an err, but got one")
	}
}

func TestFindById(t *testing.T) {

	GetInstance().tasks = append(GetInstance().tasks,
		Task{Id: 0, Description: "do dishes"},
		Task{Id: 1, Description: "do laundry"})

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
			gotTask, gotIndex, err := FindById(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTask, tt.wantTask) {
				t.Errorf("FindById() gotTask = %v, wantTask %v", gotTask, tt.wantTask)
			}
			if gotIndex != tt.wantIndex {
				t.Errorf("FindById() gotIndex = %v, wantTask %v", gotIndex, tt.wantIndex)
			}
		})
	}
}
