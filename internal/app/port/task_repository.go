package port

import (
	"fmt"
	"todo2/internal/domain"
)

type ErrTaskNotFound struct {
	Value domain.TaskID
}

type ErrStorage struct{}

func (e *ErrTaskNotFound) Error() string {
	return fmt.Sprintf("cannot find task with id %d", e.Value)
}

func (e *ErrStorage) Error() string {
	return "storage error"
}

type TaskRepository interface {
	NextID() (domain.TaskID, error)
	List() ([]*domain.Task, error)
	Get(id domain.TaskID) (*domain.Task, error)
	Save(task *domain.Task) error
	Delete(id domain.TaskID) error
}
