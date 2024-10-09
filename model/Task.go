package model

import (
	. "github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	ID          UUID       `json:"id" gorm:"primaryKey;type:char(36);"`
	Description string     `json:"description" gorm:"<-"`
	Done        bool       `json:"done" gorm:"<-"`
	UserId      UUID       `json:"user;type:char(36);"`
	CreatedAt   *time.Time `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"autoDeleteTime:true"`
}
