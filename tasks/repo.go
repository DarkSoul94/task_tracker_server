package tasks

import "github.com/DarkSoul94/task_tracker_server/models"

type TasksRepo interface {
	CreateTask(task models.Task) error
	GetTasksList(key string, user models.User) ([]*models.Task, error)
	GetTask(taskID uint64) (models.Task, error)

	InsertTaskTrack(tackTrack models.TaskTrack) error
	GetLastTaskTrack(taskID, userID uint64) (models.TaskTrack, error)

	GetCategoryByID(id uint64) (*models.Category, error)
	GetCategoryList() ([]*models.Category, error)

	GetProjectByID(id uint64) (*models.Project, error)

	Close() error
}
