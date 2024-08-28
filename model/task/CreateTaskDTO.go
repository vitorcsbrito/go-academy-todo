package task

import "github.com/google/uuid"

type CreateTaskDTO struct {
	Description string    `json:"description" binding:"required"`
	UserId      uuid.UUID `json:"userId" binding:"required"`
}
