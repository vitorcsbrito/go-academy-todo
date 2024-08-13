package model

import (
	. "github.com/google/uuid"
	"time"
)

type User struct {
	Id        UUID       `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Username  string     `json:"username" gorm:"<-"`
	Password  string     `json:"password" gorm:"<-"`
	CreatedAt *time.Time `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
}

type CreateUserDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
