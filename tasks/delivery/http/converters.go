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

func (h *Handler) toOutTask(mTask *models.Task) outTask {
	out := outTask{
		ID:           mTask.ID,
		Name:         mTask.Name,
		Description:  mTask.Description,
		InWorkTime:   mTask.InWorkTime,
		CreationDate: mTask.CreationDate.Local().String(),
		Author: &userInTask{
			ID:   mTask.Author.ID,
			Name: mTask.Author.Name,
		},
		Category: &hCategory{
			ID:   mTask.Category.ID,
			Name: mTask.Category.Name,
		},
		Status: &hStatus{
			ID:   mTask.Status.ID,
			Name: mTask.Status.Name,
		},
		Priority:  mTask.Priority,
		ExecOrder: mTask.ExecOrder,
	}

	if mTask.Developer != nil {
		out.Developer = &userInTask{
			ID:   mTask.Developer.ID,
			Name: mTask.Developer.Name,
		}
	}

	if mTask.Customer != nil {
		out.Customer = &userInTask{
			ID:   mTask.Customer.ID,
			Name: mTask.Customer.Name,
		}
	}

	if mTask.Project != nil {
		out.Project = &hProject{
			ID:   mTask.Project.ID,
			Name: mTask.Project.Name,
		}
	}

	return out
}

func (h *Handler) toOutCategory(cat *models.Category) hCategory {
	return hCategory{
		ID:   cat.ID,
		Name: cat.Name,
	}
}
