package tasks

import "github.com/DarkSoul94/task_tracker_server/models"

type TasksUC interface {
	CreateTask(task models.Task) error
	GetTasksList(user *models.User) ([]*models.TaskForList, error)
	GetTask(taskID uint64) (*models.Task, error)

	TrackTask(taskId uint64, user *models.User, status bool) error

	GetCategoryList() ([]*models.Category, error)
}
