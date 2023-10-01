package models

import (
	"time"

	"task-manager/app/domain/entities"
)

type TaskModel struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	DueDate     time.Time `db:"due_date"`
	Status      bool      `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (tm *TaskModel) ToEntities() (*entities.Task, error) {
	task, err := entities.NewTask(
		tm.ID,
		tm.Name,
		tm.Description,
		tm.DueDate,
		tm.Status,
	)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func NewTaskModelFromEntities(task *entities.Task) *TaskModel {
	return &TaskModel{
		ID:          task.GetID(),
		Name:        task.GetName(),
		Description: task.GetDescription(),
		DueDate:     task.GetDueDate(),
		Status:      task.GetStatus(),
	}
}
