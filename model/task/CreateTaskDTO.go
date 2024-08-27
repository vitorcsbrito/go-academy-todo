package task

type CreateTaskDTO struct {
	Description string `json:"description" binding:"required"`
}
