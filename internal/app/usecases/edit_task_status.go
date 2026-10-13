package usecases

import (
	"errors"
	"fmt"

	"github.com/2ips/todo/internal/app/port"
	"github.com/2ips/todo/internal/domain"
)

type EditTaskStatus struct {
	repository port.TaskRepository
}

var ErrEditTaskStatus = errors.New("edit task status error")

func NewEditTaskStatus(repository port.TaskRepository) *EditTaskStatus {
	return &EditTaskStatus{repository: repository}
}

func (uc *EditTaskStatus) Execute(id int64, status string) (*domain.Task, error) {
	taskID, err := domain.NewTaskID(id)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrEditTaskStatus, err)
	}

	newStatus, err := domain.NewTaskStatus(status)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrEditTaskStatus, err)
	}

	task, err := uc.repository.Get(taskID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrEditTaskStatus, err)
	}

	if !task.Status.IsEditable() {
		return nil, fmt.Errorf("%w; %w", ErrEditTaskStatus, &domain.ErrTaskNotEditable{Value: newStatus})
	}

	if task.Status == newStatus {
		return nil, fmt.Errorf("%w; %w", ErrEditTaskStatus, &domain.ErrTaskStatusAlreadySet{Value: task.Status})
	}

	task.Status = newStatus

	if err := uc.repository.Save(task); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrEditTaskStatus, err)
	}

	return task, nil
}
