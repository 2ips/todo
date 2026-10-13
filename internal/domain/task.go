package domain

import "time"

type Task struct {
	ID TaskID
	Title  TaskTitle
	Status TaskStatus
	Deadline TaskDeadline
}

func NewTask(id int64, title string, status string, deadline time.Time) (*Task, error) {
	task, err := ReconstructTask(id, title, status, deadline)
	if err != nil {
		return nil, err
	}

	taskDeadline, err := NewTaskDeadline(deadline)
	if err != nil {
		return nil, err
	}

	task.Deadline = taskDeadline
	return task, nil
}

func ReconstructTask(id int64, title string, status string, deadline time.Time) (*Task, error) {
	taskID, err := NewTaskID(id)
	if err != nil {
		return nil, err
	}

	taskTitle, err := NewTaskTitle(title)
	if err != nil {
		return nil, err
	}

	taskStatus, err := NewTaskStatus(status)
	if err != nil {
		return nil, err
	}

	taskDeadline, err := newTaskDeadline(deadline)
	if err != nil {
		return nil, err
	}

	return &Task{
		ID:   taskID,
		Title:    taskTitle,
		Status:   taskStatus,
		Deadline: taskDeadline,
	}, nil
}
