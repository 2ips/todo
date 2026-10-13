package port

import (
	"errors"
	"fmt"

	"github.com/2ips/todo/internal/domain"
)

type ErrTaskNotFound struct {
	Value domain.TaskID
}

var ErrStorage = errors.New("storage error")

func (e *ErrTaskNotFound) Error() string {
	return fmt.Sprintf("cannot find task with id %d", e.Value)
}

type TaskRepository interface {
	NextID() (domain.TaskID, error)
	List() ([]*domain.Task, error)
	Get(id domain.TaskID) (*domain.Task, error)
	Save(task *domain.Task) error
	Delete(id domain.TaskID) error
}
