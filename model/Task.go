package model

import (
	. "github.com/google/uuid"
	"time"
)

type Task struct {
	Id          UUID       `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Description string     `json:"description" gorm:"<-"`
	Done        bool       `json:"done" gorm:"<-"`
	CreatedAt   *time.Time `json:"created_at" gorm:"autoCreateTime:true"`
	//CreatedAt   string `json:"created_at" gorm:"autoCreateTime:false"`
}
