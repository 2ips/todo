package presentation

import (
	"fmt"
	"strconv"
	"time"

	"github.com/2ips/todo/internal/domain"
)

const TimeLayout = "DD.MM.YYYY HH:MM"

func FormatTask(task *domain.Task) string {
	return fmt.Sprintf(
		"#%d [%s] %s - due %s",
		task.ID,
		task.Status,
		task.Title,
		time.Time(task.Deadline).Format(TimeLayout),
	)
}

func ParseTaskID(raw string) (int64, error) {
	id, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("task id must be an integer, got %s", raw)
	}

	return id, nil
}

func ParseDeadline(raw string) (time.Time, error) {
	deadline, err := time.ParseInLocation(TimeLayout, raw, time.Local)
	if err != nil {
		return time.Time{}, fmt.Errorf("deadline must be %s, got %s", TimeLayout, raw)
	}

	return deadline, nil
}
