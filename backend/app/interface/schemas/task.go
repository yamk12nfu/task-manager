package schemas

import (
	"time"

	"task-manager/app/domain/entities"
)

type TaskSchema struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      bool      `json:"status"`
}

func NewTaskSchemaFromEntities(task *entities.Task) *TaskSchema {
	return &TaskSchema{
		ID:          task.GetID(),
		Name:        task.GetName(),
		Description: task.GetDescription(),
		DueDate:     task.GetDueDate(),
		Status:      task.GetStatus(),
	}
}
