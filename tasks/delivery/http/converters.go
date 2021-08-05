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

func (h *Handler) toOutTask(mTask models.Task) outTask {
	return outTask{
		ID:          mTask.ID,
		Name:        mTask.Name,
		Description: mTask.Description,
	}
}
