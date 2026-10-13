package commands

import (
	"fmt"

	"github.com/2ips/todo/internal/app/usecases"
	"github.com/2ips/todo/internal/presentation"
	"github.com/spf13/cobra"
)

type createCommand struct {
	uc *usecases.CreateTask
}

func NewCreateCommand(uc *usecases.CreateTask) *cobra.Command {
	command := &createCommand{uc: uc}

	cobraCommand := &cobra.Command{
		Use:     "create <title> <deadline>",
		Short:   "create a new task",
		Example: "create \"get an internship at Ozon\" \"13.12.2026 13:55\"",
		Args:    cobra.ExactArgs(2),
		RunE:    command.run,
	}

	return cobraCommand
}

func (c *createCommand) run(_ *cobra.Command, args []string) error {
	deadline, err := presentation.ParseDeadline(args[1])
	if err != nil {
		return err
	}

	task, err := c.uc.Execute(args[0], deadline)
	if err != nil {
		return err
	}

	fmt.Println(presentation.FormatTask(task))
	return nil
}
