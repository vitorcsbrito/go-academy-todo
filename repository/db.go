package repository

import (
	"fmt"
	"github.com/glebarez/sqlite"
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
			log.Println("Creating single instance now.")
			singleInstance = &Repository{}
		}
	}
	log.Println("Repository instance already created.")

	return singleInstance
}

func (s *Repository) Init(dialector gorm.Dialector) {
	db, err := gorm.Open(dialector, &gorm.Config{})
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	s.DB = db
}

func GetDbConnection() gorm.Dialector {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db := os.Getenv("DB")

	if db == "mysql" {
		username := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		hostname := os.Getenv("DB_HOST")
		dbname := os.Getenv("DB_NAME")

		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, hostname, dbname)
		return mysql.Open(dsn)
	} else if db == "sqlite" {
		dbname := os.Getenv("DB_NAME")
		dsn := fmt.Sprintf("file:%v.db?cache=shared", dbname)
		return sqlite.Open(dsn)
	}
	panic("DB environment variable not valid.")
}
