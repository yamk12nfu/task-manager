package database

import (
	"context"
	"errors"

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

func (r *TaskRepository) FindByID(ctx context.Context, id string) (*entities.Task, error) {
	query := "select * from tasks where id = :id"

	row, err := r.NamedQuery(ctx, query, map[string]any{"id": id})
	if err != nil {
		return &entities.Task{}, err
	}

	defer row.Close()

	tm := models.TaskModel{}
	if row.Next() {
		err = row.StructScan(&tm)
	} else {
		err = errors.New("task not found id: " + id)
	}
	if err != nil {
		return &entities.Task{}, err
	}

	task, err := tm.ToEntities()
	if err != nil {
		return &entities.Task{}, err
	}

	return task, err
}
