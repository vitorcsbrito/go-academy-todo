package model

import (
	. "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        UUID       `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Username  string     `json:"username" gorm:"<-"`
	Password  string     `json:"password" gorm:"<-"`
	Email     string     `json:"email" gorm:"<-false"`
	CreatedAt *time.Time `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"autoDeleteTime:true"`
	Tasks     []Task     `gorm:"foreignKey:UserId;references:ID"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
