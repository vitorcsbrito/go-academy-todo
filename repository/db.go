package repository

import (
	"fmt"
	"github.com/joho/godotenv"
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
