package repository

import (
	"errors"
	"fmt"
	. "go-todo-app/model"
	"sort"
	"sync"
)

var lock = &sync.Mutex{}

type Single struct {
	Tasks []Task
}

var singleInstance *Single

func GetInstance() *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &Single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

func FindById(id int) (Task, int, error) {
	for i, todo := range GetInstance().Tasks {
		if todo.Id == id {
			return todo, i, nil
		}
	}
	return Task{Id: -1}, -1, errors.New("task not found")
}

func Delete(s int) []Task {
	lastIndex := len(GetInstance().Tasks) - 1

	if len(GetInstance().Tasks) > 0 && s == lastIndex {
		GetInstance().Tasks = GetInstance().Tasks[:len(GetInstance().Tasks)-1]
		return GetInstance().Tasks
	} else {
		return append(GetInstance().Tasks[:s], GetInstance().Tasks[s+1:]...)
	}
}

func FindLatestId() int {
	if len(GetInstance().Tasks) == 0 {
		return 0
	}

	sort.SliceStable(GetInstance().Tasks, func(i, j int) bool {
		return GetInstance().Tasks[i].Id > GetInstance().Tasks[j].Id
	})

	return GetInstance().Tasks[0].Id + 1
}
