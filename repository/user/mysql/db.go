package mysql

import (
	"github.com/google/uuid"
	. "github.com/vitorcsbrito/go-academy-todo/model"
	. "github.com/vitorcsbrito/go-academy-todo/repository"
)

type MySqlRepository struct {
	*Repository
}

func NewMySqlRepository(repository *Repository) *MySqlRepository {
	repository.Init(GetMySQLConnection())

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
