package commands

import (
	"fmt"
	"strings"

	"github.com/2ips/todo/internal/app/usecases"
	"github.com/2ips/todo/internal/domain"
	"github.com/2ips/todo/internal/presentation"
	"github.com/spf13/cobra"
)

type editStatusCommand struct {
	uc *usecases.EditTaskStatus
}

func NewEditStatusCommand(uc *usecases.EditTaskStatus) *cobra.Command {
	command := &editStatusCommand{uc: uc}

	cobraCommand := &cobra.Command{
		Use:     "status <id> <status>",
		Short:   "edit task status",
		Long:    fmt.Sprintf("available task statuses: %s", strings.Join(domain.TaskStatuses, ", ")),
		Example: "status 4 in_progress",
		Args:    cobra.ExactArgs(2),
		RunE:    command.run,
	}

	return cobraCommand
}

func (c *editStatusCommand) run(_ *cobra.Command, args []string) error {
	id, err := presentation.ParseTaskID(args[0])
	if err != nil {
		return err
	}

	task, err := c.uc.Execute(id, args[1])
	if err != nil {
		return err
	}

	fmt.Println(presentation.FormatTask(task))
	return nil
}
