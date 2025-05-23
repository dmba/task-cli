package commands

import (
	"context"
	"errors"
	"github.com/urfave/cli/v3"
)

var (
	ErrRemoveTaskIdIsRequired = errors.New("task id is required")
)

type removeArgs struct {
	id int
}

func (a *removeArgs) validate() error {
	if a.id == 0 {
		return ErrRemoveTaskIdIsRequired
	}
	return nil
}

func remove(handler TaskCliCommandHandler) *cli.Command {
	var args removeArgs
	return &cli.Command{
		Name:  "delete",
		Usage: "Delete task",
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
			return handler.Delete(ctx, args.id)
		},
	}
}
