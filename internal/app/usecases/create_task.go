package usecases

import (
	"errors"
	"fmt"
	"time"

	"github.com/2ips/todo/internal/app/port"
	"github.com/2ips/todo/internal/domain"
)

type CreateTask struct {
	repository port.TaskRepository
}

var ErrCreateTask = errors.New("create task error")

func NewCreateTask(repository port.TaskRepository) *CreateTask {
	return &CreateTask{repository: repository}
}

func (uc *CreateTask) Execute(title string, deadline time.Time) (*domain.Task, error) {
	id, err := uc.repository.NextID()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCreateTask, err)
	}

	task, err := domain.NewTask(int64(id), title, string(domain.TaskPlanned), deadline)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCreateTask, err)
	}

	if err := uc.repository.Save(task); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCreateTask, err)
	}

	return task, nil
}
