package repository

import (
	"fmt"
	. "go-todo-app/errors"
	"go-todo-app/files"
	. "go-todo-app/model"
	"sort"
	"sync"
)

var lock = &sync.Mutex{}

type Repository struct {
	Tasks    *[]Task
	Filename string
	sync.Mutex
}

type InterfaceRepository interface {
	Save(task Task) int
	Update(id int, task Task) (int, error)
	FindById(id int) (*Task, int, error)
	Delete(taskId int) error
	FindAll() []Task
}

var singleInstance *Repository

func GetInstance(filename string) *Repository {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &Repository{
				Filename: filename,
			}
			singleInstance.init()
		}
	}
	fmt.Println("Repository instance already created.")

	return singleInstance
}

func (s *Repository) init() {
	tasksFromJson := files.ReadTasksFromJson(s.Filename)

	GetInstance(s.Filename).Tasks = &tasksFromJson
}

func (s *Repository) Save(task Task) int {
	s.Lock()
	defer s.Unlock()
	task.Id = s.findLatestId()

	*s.Tasks = append(*s.Tasks, task)
	files.WriteTasksToJsonFile(s.Filename, *s.Tasks)

	return task.Id
}

func (s *Repository) Update(id int, task Task) (i int, err error) {
	s.Lock()
	defer s.Unlock()

	t, i, err := s.FindById(id)

	if err != nil {
		return i, err
	}

	t.Description = task.Description
	t.Done = task.Done

	(*s.Tasks)[i] = *t

	files.WriteTasksToJsonFile(s.Filename, *s.Tasks)

	return i, nil
}

func (s *Repository) FindById(id int) (*Task, int, error) {
	for i, todo := range *s.Tasks {
		if todo.Id == id {
			return &todo, i, nil
		}
	}

	return &Task{Id: -1}, -1, NewErrTaskNotFound(id)
}

func (s *Repository) FindAll() []Task {
	return *s.Tasks
}

func (s *Repository) Delete(taskId int) error {
	s.Lock()
	defer s.Unlock()

	_, i, err := s.FindById(taskId)

	if err != nil {
		return err
	}

	tasks := *s.Tasks
	lastIndex := len(tasks) - 1

	if len(tasks) > 0 && i == lastIndex {
		tasks = tasks[:len(tasks)-1]
	} else {
		tasks = append(tasks[:i], tasks[i+1:]...)
	}

	*s.Tasks = tasks
	files.WriteTasksToJsonFile(s.Filename, *s.Tasks)

	return nil
}

func (s *Repository) findLatestId() int {
	tasks := *s.Tasks

	if len(tasks) == 0 {
		return 0
	}

	sort.SliceStable(tasks, func(i, j int) bool {
		return tasks[i].Id > tasks[j].Id
	})

	return tasks[0].Id + 1
}
