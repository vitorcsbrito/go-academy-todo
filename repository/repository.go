package repository

import (
	"errors"
	. "go-todo-app/model"
	"sort"
)

var Tasks []Task

func FindById(id int) (Task, int, error) {
	for i, todo := range Tasks {
		if todo.Id == id {
			return todo, i, nil
		}
	}
	return Task{Id: -1}, -1, errors.New("task not found")
}

func Delete(slice []Task, s int) []Task {
	lastIndex := len(slice) - 1

	if len(slice) > 0 && s == lastIndex {
		slice = slice[:len(slice)-1]
		return slice
	} else {
		return append(slice[:s], slice[s+1:]...)
	}
}

func FindLatestId() int {
	if len(Tasks) == 0 {
		return 0
	}

	sort.SliceStable(Tasks, func(i, j int) bool {
		return Tasks[i].Id > Tasks[j].Id
	})

	return Tasks[0].Id + 1
}
