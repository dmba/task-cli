package app

import (
	"context"
	"fmt"
	"github.com/dmba/task-cli/pkg/models"
	"os"
	"text/tabwriter"
	"time"
)

type Controller struct {
	tasksService *TasksService
}

func New(tasksService *TasksService) *Controller {
	return &Controller{
		tasksService: tasksService,
	}
}

func (c *Controller) Add(_ context.Context, description string) error {
	if t, err := c.tasksService.Add(description); err == nil {
		fmt.Printf("Task added successfully (ID: %d)\n", t.ID)
		return nil
	} else {
		return err
	}
}

func (c *Controller) Update(_ context.Context, id int, description string) error {
	if _, err := c.tasksService.Update(id, description); err != nil {
		return err
	}

	fmt.Printf("Task updated successfully (ID: %d)\n", id)
	return nil
}

func (c *Controller) Delete(_ context.Context, id int) error {
	if err := c.tasksService.Delete(id); err != nil {
		return err
	}

	fmt.Printf("Task deleted successfully (ID: %d)\n", id)
	return nil
}

func (c *Controller) MarkTodo(_ context.Context, id int) error {
	if _, err := c.tasksService.MarkTodo(id); err != nil {
		return err
	}

	fmt.Printf("Task marked \"todo\" successfully (ID: %d)\n", id)
	return nil
}

func (c *Controller) MarkInProgress(_ context.Context, id int) error {
	if _, err := c.tasksService.MarkInProgress(id); err != nil {
		return err
	}

	fmt.Printf("Task marked \"in-progress\" successfully (ID: %d)\n", id)
	return nil
}

func (c *Controller) MarkDone(_ context.Context, id int) error {
	if _, err := c.tasksService.MarkDone(id); err != nil {
		return err
	}

	fmt.Printf("Task marked \"done\" successfully (ID: %d)\n", id)
	return nil
}

func (c *Controller) List(_ context.Context, status models.Status) error {
	if tasks, err := c.listTasks(status); err == nil {
		return printTasks(tasks)
	} else {
		return err
	}
}

func (c *Controller) listTasks(status models.Status) ([]models.Task, error) {
	switch status {
	case models.ToDo, models.InProgress, models.Done:
		return c.tasksService.List(status)
	default:
		return c.tasksService.ListAll()
	}
}

func printTasks(tasks []models.Task) error {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	if _, err := fmt.Fprintln(w, "\tID\tDescription\tStatus\tUpdated\tCreated"); err != nil {
		return err
	}

	for i, task := range tasks {
		_, err := fmt.Fprintf(
			w,
			"%d\t%d\t%s\t%s\t%s\t%s\t\n",
			i,
			task.ID,
			task.Description,
			task.Status,
			task.UpdatedAt.Format(time.Stamp),
			task.CreatedAt.Format(time.Stamp),
		)
		if err != nil {
			return err
		}
	}

	if err := w.Flush(); err != nil {
		return err
	}
	return nil
}
