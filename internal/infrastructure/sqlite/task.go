package sqlite

import (
	"strings"
	"time"

	"github.com/2ips/todo/internal/domain"
)

type taskModel struct {
	ID         int64     `gorm:"primaryKey;autoIncrement"`
	Title      string    `gorm:"not null;size:255;check:LENGTH(title) <= 255"`
	TitleLower string    `gorm:"not null;size:255;check:LENGTH(title_lower) <= 255 AND title_lower = LOWER(title_lower)"`
	Status     string    `gorm:"not null"`
	Deadline   time.Time `gorm:"not null"`
}

func (taskModel) TableName() string {
	return "tasks"
}

func toModel(task *domain.Task) *taskModel {
	return &taskModel{
		ID:         int64(task.ID),
		Title:      string(task.Title),
		TitleLower: strings.ToLower(string(task.Title)),
		Status:     string(task.Status),
		Deadline:   time.Time(task.Deadline),
	}
}

func fromModel(task *taskModel) (*domain.Task, error) {
	return domain.ReconstructTask(task.ID, task.Title, task.Status, task.Deadline)
}
