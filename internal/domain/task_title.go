package domain

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const MaxTaskTitleLength = 255

type TaskTitle string

type ErrTaskTitleEmpty struct {
	Value string
}

type ErrTaskTitleTooLong struct {
	Length int
}

func (e ErrTaskTitleEmpty) Error() string {
	return fmt.Sprintf("task title must be non-empty, got %s", e.Value)
}

func (e ErrTaskTitleTooLong) Error() string {
	return fmt.Sprintf("task title must be shorter than %d characters, got %d", MaxTaskTitleLength, e.Length)
}

func NewTaskTitle(title string) (TaskTitle, error) {
	title = strings.TrimSpace(title)

	if title == "" {
		return "", &ErrTaskTitleEmpty{Value: title}
	}

	length := utf8.RuneCountInString(title)
	if length > MaxTaskTitleLength {
		return "", &ErrTaskTitleTooLong{Length: length}
	}

	return TaskTitle(title), nil
}
