package mysql

import (
	"database/sql"
	"fmt"

	"github.com/DarkSoul94/task_tracker_server/models"
	"github.com/DarkSoul94/task_tracker_server/pkg/logger"
	"github.com/DarkSoul94/task_tracker_server/tasks"
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
		logger.LogError("Failed insert new task to db", "task/repo/mysql", errData, err)
		return err
	}
	return nil
}

func (r *Repo) GetTasksList(key string, user models.User) ([]models.Task, error) {
	var (
		query string
	)

	args := make([]interface{}, 0)
	dbTasks := make([]dbTask, 0)
	mTasks := make([]models.Task, 0)

	switch key {
	case tasks.TargetAction_GetAllTasks:
		query = `
			SELECT * FROM tasks`
	case tasks.TargetAction_GetTaskByAuthor:
		query = `
			SELECT * FROM tasks
			WHERE author_id = ?`
		args = append(args, user.ID)
	case tasks.TargetAction_GetTaskByDev:
		query = `
			SELECT * FROM tasks
			WHERE developer = ?`
		args = append(args, user.ID)
	}
	err := r.db.Select(&dbTasks, query, args...)
	if err != nil {
		logger.LogError("Failed get tasks list from db", "task/repo/mysql", "", err)
		return nil, err
	}
	if len(dbTasks) != 0 {
		for _, val := range dbTasks {
			mTasks = append(mTasks, r.toModelTask(val))
		}
	}
	return mTasks, nil
}

func (r *Repo) Close() error {
	r.db.Close()
	return nil
}
