package commands

import (
	"context"
	"errors"
	"github.com/urfave/cli/v3"
)

var (
	ErrMarkTodoTaskIdIsRequired = errors.New("task id is required")
)

type markTodoArgs struct {
	id int
}

func (a *markTodoArgs) validate() error {
	if a.id == 0 {
		return ErrMarkTodoTaskIdIsRequired
	}
	return nil
}

func markTodo(handler TaskCliCommandHandler) *cli.Command {
	var args markTodoArgs
	return &cli.Command{
		Name:  "mark-todo",
		Usage: "Mark a task as todo",
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
			return handler.MarkTodo(ctx, args.id)
		},
	}
}
