package models

type Status string

const (
	ToDo       Status = "todo"
	InProgress Status = "in-progress"
	Done       Status = "done"
)
