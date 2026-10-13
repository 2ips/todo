package commands

import (
	"fmt"

	"github.com/2ips/todo/internal/app/usecases"
	"github.com/2ips/todo/internal/presentation"
	"github.com/spf13/cobra"
)

type listCommand struct {
	uc *usecases.ListTasks
}

func NewListCommand(uc *usecases.ListTasks) *cobra.Command {
	command := &listCommand{uc: uc}

	cobraCommand := &cobra.Command{
		Use:   "list",
		Short: "list all tasks",
		Args:  cobra.NoArgs,
		RunE:  command.run,
	}

	return cobraCommand
}

func (c *listCommand) run(*cobra.Command, []string) error {
	tasks, err := c.uc.Execute()
	if err != nil {
		return err
	}

	for _, task := range tasks {
		fmt.Println(presentation.FormatTask(task))
	}

	return nil
}
