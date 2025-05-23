package main

import (
	"github.com/dmba/task-cli/internal/app"
	"github.com/dmba/task-cli/internal/commands"
	"log"
)

const (
	outputFileName = ".task-cli.json"
)

func main() {
	controller := app.New(
		app.NewTasksService(outputFileName),
	)

	cmd := commands.NewTaskCliCommand(controller)

	if err := cmd.Exec(); err != nil {
		log.Fatal(err)
	}
}
