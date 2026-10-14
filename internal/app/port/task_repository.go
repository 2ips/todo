package port

import (
	"errors"
	"fmt"

	"github.com/2ips/todo/internal/domain"
)

type ErrTaskNotFound struct {
	Value domain.TaskID
}

type ErrInvalidPage struct {
	Value int64
}

type ErrInvalidPageLimit struct {
	Value int64
}

var ErrStorage = errors.New("storage error")

func (e *ErrTaskNotFound) Error() string {
	return fmt.Sprintf("cannot find task with id %d", e.Value)
}

func (e *ErrInvalidPage) Error() string {
	return fmt.Sprintf("page must be positive, got %d", e.Value)
}

func (e *ErrInvalidPageLimit) Error() string {
	return fmt.Sprintf("page limit must be positive, got %d", e.Value)
}

type Pagination struct {
	Page  int64
	Limit int64
}

func NewPagination(page, limit int64) (Pagination, error) {
	if page < 1 {
		return Pagination{}, &ErrInvalidPage{Value: page}
	}

	if limit < 1 {
		return Pagination{}, &ErrInvalidPageLimit{Value: limit}
	}

	return Pagination{Page: page, Limit: limit}, nil
}

func (p Pagination) Offset() int64 {
	return (p.Page - 1) * p.Limit
}

func (p Pagination) Pages(total int64) int64 {
	return (total + p.Limit - 1) / p.Limit
}

type TaskRepository interface {
	NextID() (domain.TaskID, error)
	List(title *string, pagination Pagination) ([]*domain.Task, int64, error)
	Get(id domain.TaskID) (*domain.Task, error)
	Save(task *domain.Task) error
	Delete(id domain.TaskID) error
}
