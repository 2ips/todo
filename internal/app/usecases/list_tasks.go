package usecases

import (
	"errors"
	"fmt"

	"github.com/2ips/todo/internal/app/port"
	"github.com/2ips/todo/internal/domain"
)

type ListTasks struct {
	repository port.TaskRepository
}

var ErrListTasks = errors.New("list tasks error")

func NewListTasks(repository port.TaskRepository) *ListTasks {
	return &ListTasks{repository: repository}
}

func (uc *ListTasks) Execute() ([]*domain.Task, error) {
	tasks, err := uc.repository.List()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrListTasks, err)
	}

	return tasks, nil
}
