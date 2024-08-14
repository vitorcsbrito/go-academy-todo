package service

import . "model"

func newEntity(description string) Task {
	return Task{Description: description, Done: false}
}

//
//func fromEntityToModel(te smt.TaskEntity) *Task {
//	return &Task{
//		Id:          te.Id,
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
