package usecases

import (
	"errors"
	"fmt"

	"github.com/2ips/todo/internal/app/port"
	"github.com/2ips/todo/internal/domain"
)

type DeleteTask struct {
	repository port.TaskRepository
}

var ErrDeleteTask = errors.New("delete task error")

func NewDeleteTask(repository port.TaskRepository) *DeleteTask {
	return &DeleteTask{repository: repository}
}

func (uc *DeleteTask) Execute(id int64) error {
	taskID, err := domain.NewTaskID(id)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrDeleteTask, err)
	}

	if err := uc.repository.Delete(taskID); err != nil {
		return fmt.Errorf("%w: %w", ErrDeleteTask, err)
	}

	return nil
}
