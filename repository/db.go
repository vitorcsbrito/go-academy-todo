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

func (s *Repository) Init(dl gorm.Dialector) {
	db, err := gorm.Open(dl, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	s.DB = db
}

func GetDbConnection() gorm.Dialector {

	hasEnvFile := true
	err := godotenv.Load(".env")
	if err != nil {
		hasEnvFile = false
		//log.Fatal(err)
	}

	if !hasEnvFile {
		dsn := fmt.Sprintf("file::memory:?cache=shared")
		return sqlite.Open(dsn)
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
		dsn := fmt.Sprintf("file:%v?cache=shared", dbname)
		return sqlite.Open(dsn)
	}
	panic("DB environment variable not valid.")
}
