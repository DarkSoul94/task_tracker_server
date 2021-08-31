package mysql

import (
	"database/sql"

	"github.com/DarkSoul94/task_tracker_server/models"
)

func (r *Repo) toDBTask(task models.Task) dbTask {
	dbTask := dbTask{
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		CreateDate:  task.CreateDate,
		InWorkTime:  task.InWorkTime,
		CategoryID:  task.Category.ID,
		AuthorID:    task.Author.ID,
		StatusID:    task.Status.ID,
	}
	if task.Developer != nil {
		dbTask.DeveloperID = sql.NullInt64{
			Int64: int64(task.Developer.ID),
			Valid: true,
		}
	}
	if task.Customer != nil {
		dbTask.CustomerID = sql.NullInt64{
			Int64: int64(task.Customer.ID),
			Valid: true,
		}
	}

	if task.Project != nil {
		dbTask.ProjectID = sql.NullInt64{
			Int64: int64(task.Project.ID),
			Valid: true,
		}
	}
	return dbTask
}

func (r *Repo) toModelTask(dbTask dbTask) models.Task {
	var ts models.TaskStatus
	mTask := models.Task{
		ID:          dbTask.ID,
		Name:        dbTask.Name,
		Description: dbTask.Description,
		CreateDate:  dbTask.CreateDate,
		InWorkTime:  dbTask.InWorkTime,
		Author: &models.User{
			ID: dbTask.AuthorID,
		},
		Category: &models.Category{
			ID: dbTask.CategoryID,
		},
		Priority:  dbTask.Priority,
		ExecOrder: dbTask.ExecOrder,
	}

	ts.SetByID(dbTask.StatusID)
	mTask.Status = &ts

	if dbTask.DeveloperID.Valid {
		mTask.Developer = &models.User{
			ID: uint64(dbTask.DeveloperID.Int64),
		}
	}

	if dbTask.CustomerID.Valid {
		mTask.Customer = &models.User{
			ID: uint64(dbTask.CustomerID.Int64),
		}
	}

	if dbTask.ProjectID.Valid {
		mTask.Project = &models.Project{
			ID: uint64(dbTask.ProjectID.Int64),
		}
	}

	return mTask
}

func (r *Repo) toDbTaskTrack(track models.TaskTrack) dbTaskTrack {
	dbTrack := dbTaskTrack{
		ID:         track.ID,
		TaskID:     track.TaskID,
		UserID:     track.Author.ID,
		StartTime:  track.StartTime,
		Difference: track.Difference,
	}

	if track.EndTime.IsZero() {
		dbTrack.EndTime.Valid = false
	} else {
		dbTrack.EndTime.Time = track.EndTime
		dbTrack.EndTime.Valid = true
	}

	return dbTrack
}

func (r *Repo) toModelTaskTrack(track dbTaskTrack) models.TaskTrack {
	mTrack := models.TaskTrack{
		ID:         track.ID,
		TaskID:     track.TaskID,
		Author:     &models.User{ID: track.UserID},
		StartTime:  track.StartTime,
		Difference: track.Difference,
	}

	if track.EndTime.Valid {
		mTrack.EndTime = track.EndTime.Time
	}

	return mTrack
}
