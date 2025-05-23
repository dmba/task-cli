package commands

import (
	"context"
	"github.com/dmba/task-cli/pkg/models"
	"github.com/urfave/cli/v3"
)

type listArgs struct {
	status string
}

func (a *listArgs) validate() error {
	return nil
}

func list(handler TaskCliCommandHandler) *cli.Command {
	var args listArgs
	return &cli.Command{
		Name:  "list",
		Usage: "List tasks (todo, in-progress, done)",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:        "status",
				Destination: &args.status,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if err := args.validate(); err != nil {
				return err
			}
			status := models.Status(args.status)
			switch status {
			case models.ToDo, models.InProgress, models.Done, "":
				return handler.List(ctx, status)
			default:
				return cli.Exit("Invalid status. Must be one of: todo, in-progress, done", 1)
			}
		},
	}
}
