package commands

import (
	"fmt"
	"time"

	"github.com/2ips/todo/internal/app/usecases"
	"github.com/2ips/todo/internal/presentation"
	"github.com/spf13/cobra"
)

type updateCommand struct {
	uc *usecases.UpdateTask
}

func NewUpdateCommand(uc *usecases.UpdateTask) *cobra.Command {
	command := &updateCommand{uc: uc}

	cobraCommand := &cobra.Command{
		Use:     "update <id> [--title <title>] [--deadline <deadline>]",
		Short:   "update a task",
		Example: "update 4 --title \"get an internship at Ozon\" --deadline \"13.12.2026 13:55\"",
		Args:    cobra.ExactArgs(1),
		RunE:    command.run,
	}

	cobraCommand.Flags().String("title", "", "new task title")
	cobraCommand.Flags().String("deadline", "", "new task deadline")

	return cobraCommand
}

func (c *updateCommand) run(command *cobra.Command, args []string) error {
	id, err := presentation.ParseTaskID(args[0])
	if err != nil {
		return err
	}

	flags := command.Flags()

	var title *string
	if flags.Changed("title") {
		value, err := flags.GetString("title")
		if err != nil {
			return err
		}

		title = &value
	}

	var deadline *time.Time
	if flags.Changed("deadline") {
		value, err := flags.GetString("deadline")
		if err != nil {
			return err
		}

		parsed, err := presentation.ParseDeadline(value)
		if err != nil {
			return err
		}

		deadline = &parsed
	}

	task, err := c.uc.Execute(id, title, deadline)
	if err != nil {
		return err
	}

	fmt.Println(presentation.FormatTask(task))
	return nil
}
