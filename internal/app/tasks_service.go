package app

import (
	"fmt"
	"github.com/dmba/task-cli/pkg/fs"
	"github.com/dmba/task-cli/pkg/models"
	"github.com/dmba/task-cli/pkg/utils"
	"slices"
	"time"
)

type TasksService struct {
	file *fs.Json[[]models.Task]
}

func NewTasksService(fileName string) *TasksService {
	file := fs.NewJson[[]models.Task](fileName)
	return &TasksService{
		file: file,
	}
}

func (s *TasksService) Add(description string) (*models.Task, error) {
	var newTask models.Task

	err := s.file.Modify(func(tasks []models.Task) ([]models.Task, error) {
		id := utils.NextId(tasks, func(t models.Task) int {
			return t.ID
		})
		newTask = models.Task{
			ID:          id,
			Description: description,
			Status:      models.ToDo,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		return append(tasks, newTask), nil
	})

	if err != nil {
		return nil, err
	}
	return &newTask, nil
}

func (s *TasksService) Update(id int, description string) (*models.Task, error) {
	return s.updateTask(id, func(t *models.Task) {
		t.Description = description
		t.UpdatedAt = time.Now()
	})
}

func (s *TasksService) Delete(id int) error {
	return s.file.Modify(func(tasks []models.Task) ([]models.Task, error) {
		modified := slices.DeleteFunc(tasks, func(e models.Task) bool {
			return e.ID == id
		})
		if len(tasks) == len(modified) {
			return nil, fmt.Errorf("task \"%d\" not found", id)
		}

		return modified, nil
	})
}

func (s *TasksService) MarkTodo(id int) (*models.Task, error) {
	return s.updateTask(id, func(t *models.Task) {
		t.Status = models.ToDo
		t.UpdatedAt = time.Now()
	})
}

func (s *TasksService) MarkInProgress(id int) (*models.Task, error) {
	return s.updateTask(id, func(t *models.Task) {
		t.Status = models.InProgress
		t.UpdatedAt = time.Now()
	})
}

func (s *TasksService) MarkDone(id int) (*models.Task, error) {
	return s.updateTask(id, func(t *models.Task) {
		t.Status = models.Done
		t.UpdatedAt = time.Now()
	})
}

func (s *TasksService) List(status models.Status) ([]models.Task, error) {
	tasks, err := s.file.Read()
	if err != nil {
		return nil, err
	}
	tasks = utils.Filter(tasks, filterByStatus(status))
	slices.SortFunc(tasks, models.ByCompleteness)
	return tasks, nil
}

func (s *TasksService) ListAll() ([]models.Task, error) {
	tasks, err := s.file.Read()
	if err != nil {
		return nil, err
	}
	slices.SortFunc(tasks, models.ByCompleteness)
	return tasks, nil
}

func (s *TasksService) updateTask(id int, update func(*models.Task)) (*models.Task, error) {
	var modifiedTask *models.Task

	err := s.file.Modify(func(tasks []models.Task) ([]models.Task, error) {
		idx := slices.IndexFunc(tasks, indexById(id))
		if idx == -1 {
			return nil, fmt.Errorf("task \"%d\" not found", id)
		}
		modifiedTask = &tasks[idx]
		update(modifiedTask)
		return tasks, nil
	})

	if err != nil {
		return nil, err
	}
	return modifiedTask, nil
}

func indexById(id int) func(models.Task) bool {
	return func(task models.Task) bool {
		return task.ID == id
	}
}

func filterByStatus(status models.Status) utils.FilterFunc[models.Task] {
	return func(task models.Task) bool {
		return task.Status == status
	}
}
