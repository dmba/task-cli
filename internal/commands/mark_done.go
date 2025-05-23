package commands

import (
	"context"
	"errors"
	"github.com/urfave/cli/v3"
)

var (
	ErrMarkDoneTaskIdIsRequired = errors.New("task id is required")
)

type markDoneArgs struct {
	id int
}

func (a *markDoneArgs) validate() error {
	if a.id == 0 {
		return ErrMarkDoneTaskIdIsRequired
	}
	return nil
}

func markDone(handler TaskCliCommandHandler) *cli.Command {
	var args markDoneArgs
	return &cli.Command{
		Name:  "mark-done",
		Usage: "Mark a task as done",
		Arguments: []cli.Argument{
			&cli.IntArg{
				Name:        "id",
				Destination: &args.id,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if err := args.validate(); err != nil {
				return err
			}
			return handler.MarkDone(ctx, args.id)
		},
	}
}
