package mapper

import (
	. "github.com/vitorcsbrito/go-academy-todo/model"
)

func NewEntity(description string, user *User) Task {
	return Task{
		Description: description,
		Done:        false,
		UserId:      user.ID,
	}
}

//
//func fromEntityToModel(te smt.TaskEntity) *Task {
//	return &Task{
//		ID:          te.ID,
//		Description: te.Description,
//		Done:        te.Done,
//		CreatedAt:   te.CreatedAt,
//	}
//}
//
//func fromEntityListToModel(te []smt.TaskEntity) []Task {
//	resArr := make([]Task, len(te))
//
//	for i, t := range te {
//		task := *fromEntityToModel(t)
//		resArr[i] = task
//	}
//
//	return resArr
//}
