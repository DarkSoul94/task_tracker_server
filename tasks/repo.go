package tasks

import "github.com/DarkSoul94/task_tracker_server/models"

type TasksRepo interface {
	CreateTask(task models.Task) error
	Close() error
}
