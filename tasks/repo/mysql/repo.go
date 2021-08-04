package mysql

import (
	"database/sql"
	"fmt"

	"github.com/DarkSoul94/task_tracker_server/models"
	"github.com/DarkSoul94/task_tracker_server/pkg/logger"
	"github.com/jmoiron/sqlx"
)

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		db: sqlx.NewDb(db, "mysql"),
	}
}

func (r *Repo) CreateTask(task models.Task) error {
	dbTask := r.toDBTask(task)
	query := `
		INSERT INTO tasks SET
		id = :id,
		name = :name,
		description = :description,
		create_time = :create_time,
		in_work_time = :in_work_time,
		author_id = :author_id,
		status_id = :status_id`
	_, err := r.db.NamedExec(query, dbTask)
	if err != nil {
		errData := fmt.Sprintf("task author: %d %s", task.Author.ID, task.Author.Name)
		logger.LogError("Failed insert new task to db", "helpdesk/repo/mysql", errData, err)
		return err
	}
	return nil
}

func (r *Repo) Close() error {
	r.db.Close()
	return nil
}
