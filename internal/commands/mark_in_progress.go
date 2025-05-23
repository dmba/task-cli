package commands

import (
	"context"
	"errors"
	"github.com/urfave/cli/v3"
)

var (
	ErrMarkInProgressTaskIdIsRequired = errors.New("task id is required")
)

type markInProgressArgs struct {
	id int
}

func (a *markInProgressArgs) validate() error {
	if a.id == 0 {
		return ErrMarkInProgressTaskIdIsRequired
	}
	return nil
}

func markInProgress(handler TaskCliCommandHandler) *cli.Command {
	var args markInProgressArgs
	return &cli.Command{
		Name:  "mark-in-progress",
		Usage: "Mark a task as in-progress",
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
			return handler.MarkInProgress(ctx, args.id)
		},
	}
}
