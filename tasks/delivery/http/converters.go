package http

import "github.com/DarkSoul94/task_tracker_server/models"

func (h *Handler) toNewModelTask(task newTask, user *models.User) models.Task {
	mTask := models.Task{
		Name:        task.Name,
		Description: task.Description,
		Author:      user,
		Category: &models.Category{
			ID: task.CategoryID,
		},
		Priority:  task.Priority,
		ExecOrder: task.ExecOrder,
	}

	if task.DeveloperID != 0 {
		mTask.Developer = &models.User{
			ID: task.DeveloperID,
		}
	}

	if task.CustomerID != 0 {
		mTask.Customer = &models.User{
			ID: task.CustomerID,
		}
	}

	if task.ProjectID != 0 {
		mTask.Project = &models.Project{
			ID: task.ProjectID,
		}
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
