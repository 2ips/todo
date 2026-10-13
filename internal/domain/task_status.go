package domain

import (
	"fmt"
	"strings"
)

type TaskStatus string

const (
	TaskPlanned    TaskStatus = "planned"
	TaskInProgress TaskStatus = "in_progress"
	TaskCancelled  TaskStatus = "cancelled"
	TaskDone       TaskStatus = "done"
)

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
	taskStatuses := []string{
		string(TaskPlanned),
		string(TaskInProgress),
		string(TaskCancelled),
		string(TaskDone),
	}

	return fmt.Sprintf("expected task status as [%s], got %s", strings.Join(taskStatuses, ", "), e.Value)
}

func (e ErrTaskStatusAlreadySet) Error() string {
	return fmt.Sprintf("task status already set to %s", e.Value)
}

func (e ErrTaskNotEditable) Error() string {
	return fmt.Sprintf("cannot edit task with status %s", e.Value)
}

func NewTaskStatus(status string) (TaskStatus, error) {
	taskStatus := TaskStatus(strings.ToLower(status))
	switch taskStatus {
	case TaskPlanned, TaskInProgress, TaskCancelled, TaskDone:
		return taskStatus, nil
	default:
		return "", &ErrUnknownTaskStatus{Value: status}
	}
}

func (s TaskStatus) IsEditable() bool {
	return s != TaskDone && s != TaskCancelled
}
