package tasks

import "github.com/DarkSoul94/task_tracker_server/models"

type TasksUC interface {
	GetTasksList(userID uint64) ([]models.Task, error)
}
