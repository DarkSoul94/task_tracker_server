package http

import "github.com/DarkSoul94/task_tracker_server/models"

func (h *Handler) toNewModelTask(task newTask, user *models.User) models.Task {
	mTask := models.Task{
		Name:        task.Name,
		Description: task.Description,
		Author:      user,
	}
	return mTask
}
