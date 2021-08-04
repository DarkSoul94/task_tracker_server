package tasksUC

import "errors"

var (
	ErrFailedCreateTask   = errors.New("Failed to create new task")
	ErrFailedGetTasksList = errors.New("Failed get tasks list")
)
