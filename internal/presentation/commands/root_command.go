package commands

import (
	"github.com/2ips/todo/internal/app/usecases"
	"github.com/spf13/cobra"
)

func NewRootCommand(
	create *usecases.CreateTask,
	get *usecases.GetTask,
	update *usecases.UpdateTask,
	list *usecases.ListTasks,
	status *usecases.EditTaskStatus,
	deleteTask *usecases.DeleteTask,
) *cobra.Command {
	root := &cobra.Command{
		Use:           "todo",
		Short:         "todo manager",
		SilenceErrors: true,
	}

	root.AddCommand(
		NewCreateCommand(create),
		NewGetCommand(get),
		NewUpdateCommand(update),
		NewListCommand(list),
		NewEditStatusCommand(status),
		NewDeleteCommand(deleteTask),
	)

	return root
}
