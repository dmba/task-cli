package commands

import "github.com/urfave/cli/v3"

func root(handler TaskCliCommandHandler) *cli.Command {
	return &cli.Command{
		Name:  "task-cli",
		Usage: "Task tracker",
		Commands: []*cli.Command{
			add(handler),
			update(handler),
			remove(handler),
			markTodo(handler),
			markInProgress(handler),
			markDone(handler),
			list(handler),
		},
	}
}
