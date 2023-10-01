package database

import (
	"context"

	"task-manager/app/domain/entities"
	"task-manager/app/interface/models"
)

type TaskRepository struct {
	SQLHandler
}

func (r *TaskRepository) Save(ctx context.Context, task *entities.Task) (int64, error) {
	tm := models.NewTaskModelFromEntities(task)

	query := "insert into tasks (id, name, description, due_date, status) values (:id, :name, :description, :due_date, :status)"

	res, err := r.NamedExec(ctx, query, tm)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}
