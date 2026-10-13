package domain

import (
	"fmt"
	"slices"
	"strings"
)

type TaskStatus string

const (
	TaskPlanned    TaskStatus = "planned"
	TaskInProgress TaskStatus = "in_progress"
	TaskCancelled  TaskStatus = "cancelled"
	TaskDone       TaskStatus = "done"
)

var TaskStatuses = []string{
	string(TaskPlanned),
	string(TaskInProgress),
	string(TaskCancelled),
	string(TaskDone),
}

type ErrUnknownTaskStatus struct {
	Value string
}

type ErrTaskStatusAlreadySet struct {
	Value TaskStatus
}

type ErrTaskNotEditable struct {
	Value TaskStatus
}

func (e ErrUnknownTaskStatus) Error() string {
	return fmt.Sprintf("expected task status as [%s], got %s", strings.Join(TaskStatuses, ", "), e.Value)
}

func (e ErrTaskStatusAlreadySet) Error() string {
	return fmt.Sprintf("task status already set to %s", e.Value)
}

func (e ErrTaskNotEditable) Error() string {
	return fmt.Sprintf("cannot edit task with status %s", e.Value)
}

func NewTaskStatus(status string) (TaskStatus, error) {
	status = strings.ToLower(status)

	if slices.Contains(TaskStatuses, status) {
		return TaskStatus(status), nil
	}
	return "", ErrUnknownTaskStatus{Value: status}
}

func (s TaskStatus) IsEditable() bool {
	return s != TaskDone && s != TaskCancelled
}
