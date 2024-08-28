package model

import (
	. "github.com/google/uuid"
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
