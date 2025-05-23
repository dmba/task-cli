package models

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
	"time"
)

func TestSortsTasksByStatusThenCreatedAt(t *testing.T) {
	now := time.Now()

	actual := []Task{
		{Status: InProgress, CreatedAt: now.Add(-3 * time.Hour)},
		{Status: Done, CreatedAt: now.Add(-2 * time.Hour)},
		{Status: ToDo, CreatedAt: now.Add(-1 * time.Hour)},
		{Status: ToDo, CreatedAt: now.Add(-3 * time.Hour)},
		{Status: InProgress, CreatedAt: now},
	}
	expected := []Task{
		{Status: ToDo, CreatedAt: now.Add(-1 * time.Hour)},
		{Status: ToDo, CreatedAt: now.Add(-3 * time.Hour)},
		{Status: InProgress, CreatedAt: now},
		{Status: InProgress, CreatedAt: now.Add(-3 * time.Hour)},
		{Status: Done, CreatedAt: now.Add(-2 * time.Hour)},
	}

	slices.SortFunc(actual, ByCompleteness)

	assert.Equal(t, expected, actual)
}

func TestHandlesEmptyTaskList(t *testing.T) {
	var actual []Task
	var expected []Task

	slices.SortFunc(actual, ByCompleteness)

	assert.Equal(t, expected, actual)
}

func TestHandlesSingleTaskInList(t *testing.T) {
	now := time.Now()

	actual := []Task{
		{Status: ToDo, CreatedAt: now},
	}
	expected := []Task{
		{Status: ToDo, CreatedAt: now},
	}

	slices.SortFunc(actual, ByCompleteness)

	assert.Equal(t, expected, actual)
}
