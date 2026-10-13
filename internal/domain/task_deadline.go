package domain

import (
	"fmt"
	"time"
)

type TaskDeadline time.Time

type ErrZeroTaskDeadline struct {
	Value time.Time
}

type ErrTaskDeadlineInThePast struct {
	Value     time.Time
	Threshold time.Time
}

func (e *ErrZeroTaskDeadline) Error() string {
	return fmt.Sprintf("task deadline must be non-zero, got %s", e.Value)
}

func (e *ErrTaskDeadlineInThePast) Error() string {
	return fmt.Sprintf("task deadline must be after %s, got %s", e.Threshold, e.Value)
}

func NewTaskDeadline(deadline time.Time) (TaskDeadline, error) {
	taskDeadline, err := newTaskDeadline(deadline)
	if err != nil {
		return TaskDeadline{}, err
	}

	threshold := time.Now()
	if time.Time(taskDeadline).Before(threshold) {
		return TaskDeadline{}, &ErrTaskDeadlineInThePast{Value: deadline, Threshold: threshold}
	}

	return taskDeadline, nil
}

func newTaskDeadline(deadline time.Time) (TaskDeadline, error) {
	deadline = deadline.UTC()

	if deadline.IsZero() {
		return TaskDeadline{}, &ErrZeroTaskDeadline{Value: deadline}
	}

	return TaskDeadline(deadline), nil
}
