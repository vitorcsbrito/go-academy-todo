package testutils

import (
	"fmt"
	. "go-todo-app/model"
	. "go-todo-app/repository"
	. "go-todo-app/service"
	"os"
	"time"
)

func getTestDbFileName() string {
	return fmt.Sprintf("test_%d.json", time.Now().Unix())
}

func SetupTaskService() (taskService *TaskService, fn string) {
	testRepo, fn := SetupTaskRepository()

	taskService = NewTaskService(testRepo)
	return
}

func SetupTaskRepository() (*Repository, string) {
	filename := getTestDbFileName()
	testRepo := GetInstance(filename)

	testRepo.Save(Task{Id: 0, Description: "do dishes"})
	testRepo.Save(Task{Id: 1, Description: "do laundry"})

	return testRepo, filename
}

func Cleanup(filename string) {
	err := os.Remove(filename)

	if err != nil {
		panic(err)
	}
}
