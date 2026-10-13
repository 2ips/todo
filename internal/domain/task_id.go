package domain

import "fmt"

type TaskID int64

type ErrInvalidTaskID struct{
	Value int64
}

func (e *ErrInvalidTaskID) Error() string {
	return fmt.Sprintf("task id must be a positive, got %d", e.Value)
}

func NewTaskID(id int64) (TaskID, error) {
	if id <= 0 {
		return 0, &ErrInvalidTaskID{Value: id}
	}
	return TaskID(id), nil
}
