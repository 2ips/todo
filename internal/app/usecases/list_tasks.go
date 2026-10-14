package usecases

import (
	"errors"
	"fmt"
	"strings"

	"github.com/2ips/todo/internal/app/port"
	"github.com/2ips/todo/internal/domain"
)

const (
	defaultPage      = 1
	defaultPageLimit = 20
)

type ListTasks struct {
	repository port.TaskRepository
}

var ErrListTasks = errors.New("list tasks error")

func NewListTasks(repository port.TaskRepository) *ListTasks {
	return &ListTasks{repository: repository}
}

func (uc *ListTasks) Execute(title *string, page *int64, limit *int64) ([]*domain.Task, int64, int64, error) {
	if page == nil {
		page = new(int64(defaultPage))
	}

	if limit == nil {
		limit = new(int64(defaultPageLimit))
	}

	pagination, err := port.NewPagination(*page, *limit)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("%w: %w", ErrListTasks, err)
	}

	if title != nil {
		trimmed := strings.TrimSpace(*title)
		if trimmed == "" {
			title = nil
		} else {
			title = &trimmed
		}
	}

	tasks, total, err := uc.repository.List(title, pagination)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("%w: %w", ErrListTasks, err)
	}

	return tasks, total, pagination.Pages(total), nil
}
