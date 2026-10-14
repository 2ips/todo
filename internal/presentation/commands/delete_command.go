package commands

import (
	"fmt"

	"github.com/2ips/todo/internal/app/usecases"
	"github.com/2ips/todo/internal/presentation"
	"github.com/spf13/cobra"
)

type deleteCommand struct {
	uc *usecases.DeleteTask
}

func NewDeleteCommand(uc *usecases.DeleteTask) *cobra.Command {
	command := &deleteCommand{uc: uc}

	cobraCommand := &cobra.Command{
		Use:     "delete <id>",
		Short:   "delete a task",
		Example: "delete 4",
		Args:    cobra.ExactArgs(1),
		RunE:    command.run,
	}

	return cobraCommand
}

func (c *deleteCommand) run(_ *cobra.Command, args []string) error {
	id, err := presentation.ParseTaskID(args[0])
	if err != nil {
		return err
	}

	if err := c.uc.Execute(id); err != nil {
		return err
	}

	fmt.Printf("deleted task %d\n", id)
	return nil
}
