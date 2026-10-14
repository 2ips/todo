package usecases

import (
	"errors"
	"fmt"
	"time"

	"github.com/2ips/todo/internal/app/port"
	"github.com/2ips/todo/internal/domain"
)

type UpdateTask struct {
	repository port.TaskRepository
}

var ErrUpdateTask = errors.New("error updating task")

func NewUpdateTask(repository port.TaskRepository) *UpdateTask {
	return &UpdateTask{repository: repository}
}

func (uc *UpdateTask) Execute(id int64, title *string, deadline *time.Time) (*domain.Task, error) {
	taskID, err := domain.NewTaskID(id)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrUpdateTask, err)
	}

	task, err := uc.repository.Get(taskID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrUpdateTask, err)
	}

	if !task.Status.IsEditable() {
		return nil, fmt.Errorf("%w: %w", ErrUpdateTask, domain.ErrTaskNotEditable{Value: task.Status})
	}

	if title != nil {
		task.Title, err = domain.NewTaskTitle(*title)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrUpdateTask, err)
		}
	}

	if deadline != nil {
		task.Deadline, err = domain.NewTaskDeadline(*deadline)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrUpdateTask, err)
		}
	}

	if err := uc.repository.Save(task); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrUpdateTask, err)
	}

	return task, nil
}
