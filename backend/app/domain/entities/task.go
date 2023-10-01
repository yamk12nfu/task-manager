package entities

import (
	"errors"
	"time"
)

type Task struct {
	id          string
	name        string
	description string
	dueDate     time.Time
	status      bool
}

func NewTask(id, name, description string, dueDate time.Time, status bool) (*Task, error) {
	if name == "" || dueDate.IsZero() {
		return nil, errors.New("必須項目が入力されていません")
	}

	return &Task{
		id:          id,
		name:        name,
		description: description,
		dueDate:     dueDate,
		status:      status,
	}, nil
}

func (t *Task) GetID() string {
	return t.id
}

func (t *Task) GetName() string {
	return t.name
}

func (t *Task) GetDescription() string {
	return t.description
}

func (t *Task) GetDueDate() time.Time {
	return t.dueDate
}

func (t *Task) GetStatus() bool {
	return t.status
}

func (t *Task) Done() {
	t.status = true
}
