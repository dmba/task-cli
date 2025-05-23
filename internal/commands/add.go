package commands

import (
	"context"
	"errors"
	"github.com/urfave/cli/v3"
)

var (
	ErrAddTaskDescriptionRequired = errors.New("task description is required")
)

type addArgs struct {
	description string
}

func (a *addArgs) validate() error {
	if a.description == "" {
		return ErrAddTaskDescriptionRequired
	}
	return nil
}

func add(handler TaskCliCommandHandler) *cli.Command {
	var args addArgs
	return &cli.Command{
		Name:  "add",
		Usage: "Add task",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:        "description",
				Destination: &args.description,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if err := args.validate(); err != nil {
				return err
			}
			return handler.Add(ctx, args.description)
		},
	}
}
