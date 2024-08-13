package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	. "go-todo-app/errors"
	. "go-todo-app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
)

var lock = &sync.Mutex{}

type Repository struct {
	DB *gorm.DB
}

type InterfaceRepository interface {
	Save(task Task) uuid.UUID
	Update(id uuid.UUID, task Task) (uuid.UUID, error)
	FindById(id uuid.UUID) (*Task, uuid.UUID, error)
	Delete(taskId *Task) error
	FindAll() ([]Task, error)
}

var singleInstance *Repository

func GetInstance() *Repository {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &Repository{}
			singleInstance.init()
		}
	}
	fmt.Println("Repository instance already created.")

	return singleInstance
}

func (s *Repository) init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, hostname, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	s.DB = db
}

func (s *Repository) Save(task Task) uuid.UUID {

	//s2 := time.Now().Format(time.RFC3339)
	t := Task{Description: task.Description, Done: task.Done}
	//t := Task{Description: task.Description, Done: task.Done, CreatedAt: s2}
	newUUID, _ := uuid.NewUUID()
	t.Id = newUUID

	_ = s.DB.Save(&t)

	return t.Id
}

func (s *Repository) Update(id uuid.UUID, task Task) (i uuid.UUID, err error) {
	t, i, err := s.FindById(id)

	if err != nil {
		return i, err
	}

	t.Description = task.Description
	t.Done = task.Done

	res := s.DB.Save(&t)

	if res.Error != nil {
		return i, res.Error
	}

	return i, nil
}

func (s *Repository) FindById(id uuid.UUID) (*Task, uuid.UUID, error) {

	var foundTask Task
	res := s.DB.Find(&foundTask, id)

	newUUID, _ := uuid.NewUUID()
	if res.Error == nil {
		return &foundTask, newUUID, nil
	}

	return nil, newUUID, NewErrTaskNotFound(id)
}

func (s *Repository) FindAll() ([]Task, error) {

	var foundTask []Task

	res := s.DB.Order("created_at asc").Find(&foundTask)
	if res.Error != nil {
		return nil, res.Error
	}

	return foundTask, nil
}

func (s *Repository) Delete(task *Task) error {
	res := s.DB.Delete(task)

	return res.Error
}
