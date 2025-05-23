package commands

import (
	"context"
	"errors"
	"github.com/urfave/cli/v3"
)

var (
	ErrUpdateTaskIdIsRequired   = errors.New("task id is required")
	ErrUpdateTaskDescIsRequired = errors.New("task description is required")
)

type updateArgs struct {
	id          int
	description string
}

func (a *updateArgs) validate() error {
	if a.id == 0 {
		return ErrUpdateTaskIdIsRequired
	}
	if a.description == "" {
		return ErrUpdateTaskDescIsRequired
	}
	return nil
}

func update(handler TaskCliCommandHandler) *cli.Command {
	var args updateArgs
	return &cli.Command{
		Name:  "update",
		Usage: "Update task",
		Arguments: []cli.Argument{
			&cli.IntArg{
				Name:        "id",
				Destination: &args.id,
			},
			&cli.StringArg{
				Name:        "description",
				Destination: &args.description,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if err := args.validate(); err != nil {
				return err
			}
			return handler.Update(ctx, args.id, args.description)
		},
	}
}
