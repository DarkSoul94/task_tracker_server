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
		AuthorID:    task.Author.ID,
		StatusID:    task.Status.ID,
	}
	if task.Developer != nil {
		dbTask.Developer = sql.NullInt64{
			Int64: int64(task.Developer.ID),
			Valid: true,
		}
	}
	return dbTask
}
