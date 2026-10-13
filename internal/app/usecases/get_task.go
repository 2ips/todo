package usecases

import (
	"errors"
	"fmt"

	"github.com/2ips/todo/internal/app/port"
	"github.com/2ips/todo/internal/domain"
)

type GetTask struct {
	repository port.TaskRepository
}

var ErrGetTask = errors.New("get task error")

func NewGetTask(repository port.TaskRepository) *GetTask {
	return &GetTask{repository: repository}
}

func (uc *GetTask) Execute(id int64) (*domain.Task, error) {
	taskID, err := domain.NewTaskID(id)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetTask, err)
	}

	task, err := uc.repository.Get(taskID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetTask, err)
	}

	return task, nil
}
