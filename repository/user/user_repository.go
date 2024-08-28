package user

import (
	"github.com/google/uuid"
	. "github.com/vitorcsbrito/go-academy-todo/model"
)

type Repository interface {
	Save(task User) (uuid.UUID, error)
	Update(id uuid.UUID, task User) (uuid.UUID, error)
	Get(id uuid.UUID) (user User, err error)
	//FindById(id uuid.UUID) (*User, uuid.UUID, error)
	//Delete(userId *User) error
	//FindAll() ([]User, error)
}
