package commands

import (
	"context"
	"github.com/dmba/task-cli/pkg/models"
	"github.com/urfave/cli/v3"
	"os"
)

type TaskCliCommandHandler interface {
	Add(ctx context.Context, description string) error
	Update(ctx context.Context, id int, description string) error
	Delete(ctx context.Context, id int) error
	MarkTodo(ctx context.Context, id int) error
	MarkInProgress(ctx context.Context, id int) error
	MarkDone(ctx context.Context, id int) error
	List(ctx context.Context, status models.Status) error
}

type TaskCliCommand struct {
	*cli.Command
}

func NewTaskCliCommand(handler TaskCliCommandHandler) *TaskCliCommand {
	return &TaskCliCommand{
		Command: root(handler),
	}
}

func (tc *TaskCliCommand) Exec() error {
	return tc.Run(context.Background(), os.Args)
}
