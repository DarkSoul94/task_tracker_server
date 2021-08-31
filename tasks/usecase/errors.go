package tasksUC

import "errors"

var (
	ErrFailedCreateTask   = errors.New("Failed to create new task")
	ErrFailedGetTasksList = errors.New("Failed get tasks list")
	ErrNotDeveloper       = errors.New("You are not a developer in this task")
	ErrIsTracking         = errors.New("You are already tracking this task")
	ErrIsNotTracking      = errors.New("You are not tracking this task")
	ErrNotInWork          = errors.New("The task is not at work")
)
