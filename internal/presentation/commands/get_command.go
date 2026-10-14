package commands

import (
	"fmt"

	"github.com/2ips/todo/internal/app/usecases"
	"github.com/2ips/todo/internal/presentation"
	"github.com/spf13/cobra"
)

type getCommand struct {
	uc *usecases.GetTask
}

func NewGetCommand(uc *usecases.GetTask) *cobra.Command {
	command := &getCommand{uc: uc}

	cobraCommand := &cobra.Command{
		Use:     "get <id>",
		Short:   "show one task",
		Example: "get 4",
		Args:    cobra.ExactArgs(1),
		RunE:    command.run,
	}

	return cobraCommand
}

func (c *getCommand) run(_ *cobra.Command, args []string) error {
	id, err := presentation.ParseTaskID(args[0])
	if err != nil {
		return err
	}

	task, err := c.uc.Execute(id)
	if err != nil {
		return err
	}

	fmt.Println(presentation.FormatTask(task))
	return nil
}
