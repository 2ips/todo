package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/2ips/todo/internal/app/usecases"
	"github.com/2ips/todo/internal/infrastructure/sqlite"
	"github.com/2ips/todo/internal/presentation/commands"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func databasePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("find user home directory: %w", err)
	}

	return filepath.Join(home, "todo.db"), nil
}

func run() error {
	path, err := databasePath()
	if err != nil {
		return err
	}

	db, err := sqlite.Open(path)
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}

	repository := sqlite.NewTaskRepository(db)

	root := commands.NewRootCommand(
		usecases.NewCreateTask(repository),
		usecases.NewGetTask(repository),
		usecases.NewUpdateTask(repository),
		usecases.NewListTasks(repository),
		usecases.NewEditTaskStatus(repository),
		usecases.NewDeleteTask(repository),
	)

	return root.Execute()
}
