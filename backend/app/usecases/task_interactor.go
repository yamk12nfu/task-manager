package usecases

import (
	"context"
	"time"

	"task-manager/app/domain/entities"
)

type TaskRepository interface {
	Save(context.Context, *entities.Task) (int64, error)
}

type TaskInteractor struct {
	Transaction        Transaction
	TaskRepository     TaskRepository
	UniqueIDRepository UniqueIDRepository
}

func (i *TaskInteractor) Save(ctx context.Context, name, description string, dueDate time.Time, status bool) error {
	return i.Transaction.Do(ctx, func(ctx context.Context) (err error) {
		id, err := i.UniqueIDRepository.Issue(ctx)
		if err != nil {
			return
		}

		task, err := entities.NewTask(id, name, description, dueDate, status)
		if err != nil {
			return
		}

		_, err = i.TaskRepository.Save(ctx, task)

		return
	})

}
