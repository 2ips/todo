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
		Use:     "list [--title <title>] [--page <page>] [--limit <limit>]",
		Short:   "list all tasks",
		Example: "list --title ozon --page 4 --limit 10",
		Args:    cobra.NoArgs,
		RunE:    command.run,
	}

	cobraCommand.Flags().String("title", "", "case insensitive search by title")
	cobraCommand.Flags().Int64("page", 0, "page number")
	cobraCommand.Flags().Int64("limit", 0, "tasks per page")

	return cobraCommand
}

func (c *listCommand) run(command *cobra.Command, _ []string) error {
	flags := command.Flags()

	var title *string
	if flags.Changed("title") {
		value, err := flags.GetString("title")
		if err != nil {
			return err
		}

		title = &value
	}

	var page *int64
	if flags.Changed("page") {
		value, err := flags.GetInt64("page")
		if err != nil {
			return err
		}

		page = &value
	}

	var limit *int64
	if flags.Changed("limit") {
		value, err := flags.GetInt64("limit")
		if err != nil {
			return err
		}

		limit = &value
	}

	tasks, total, pages, err := c.uc.Execute(title, page, limit)
	if err != nil {
		return err
	}

	for _, task := range tasks {
		fmt.Println(presentation.FormatTask(task))
	}

	fmt.Printf("total: %d, pages: %d\n", total, pages)
	return nil
}
