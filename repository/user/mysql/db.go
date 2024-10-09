package mysql

import (
	"github.com/google/uuid"
	. "github.com/vitorcsbrito/go-academy-todo/model"
	. "github.com/vitorcsbrito/go-academy-todo/repository"
	"github.com/vitorcsbrito/utils/errors"
	"log"
)

type MySqlRepository struct {
	*Repository
}

func NewMySqlRepository(repository *Repository) *MySqlRepository {
	repository.Init(GetDbConnection())
	err1 := repository.DB.AutoMigrate(&User{})
	if err1 != nil {
		log.Fatal(err1)
	}

	return &MySqlRepository{repository}
}

func (s *MySqlRepository) Save(user User) (uuid.UUID, error) {
	newUUID, _ := uuid.NewUUID()
	user.ID = newUUID

	res := s.DB.Create(&user)

	return user.ID, res.Error
}

func (s *MySqlRepository) Update(id uuid.UUID, task User) (uuid.UUID, error) {

	newUUID, _ := uuid.NewUUID()
	return newUUID, nil
}

func (s *MySqlRepository) Get(id uuid.UUID) (user User, err error) {

	var foundUser User
	res := s.DB.First(&foundUser, id)

	return foundUser, res.Error
}

func (s *MySqlRepository) GetAll() (users []User, err error) {

	var foundUsers []User
	res := s.DB.Find(&foundUsers)

	return foundUsers, res.Error
}

func (s *MySqlRepository) GetByUsername(username string) (user User, err error) {
	var foundUser User
	res := s.DB.Where("username = ?", username).First(&foundUser)

	if res.RowsAffected < 1 {
		return User{}, errors.ErrUserNotFound
	}

	return foundUser, res.Error
}
