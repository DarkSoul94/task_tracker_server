package mysql

import (
	"database/sql"
	"fmt"
	"strconv"

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
		creation_date = :creation_date,
		in_work_time = :in_work_time,
		author_id = :author_id,
		developer_id = :developer_id,
		customer_id = :customer_id,
		status_id = :status_id,
		category_id = :category_id,
		project_id = :project_id,
		priority = :priority,
		exec_order = :exec_order`
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
	case tasks.KeyGet_All:
		query = `
			SELECT * FROM tasks`
	case tasks.KeyGet_Author:
		query = `
			SELECT * FROM tasks
			WHERE author_id = ?`
		args = []interface{}{user.ID}
	case tasks.KeyGet_Dev:
		query = `
			SELECT * FROM tasks
			WHERE developer_id = ?`
		args = []interface{}{user.ID}
	case tasks.KeyGet_AuthorDev:
		query = `
		SELECT * FROM tasks
		WHERE developer_id = ?
		OR author_id = ?`
		args = []interface{}{user.ID, user.ID}
	case tasks.KeyGet_Customer:
		query = `
		SELECT * FROM tasks
		WHERE customer_id = ?`
		args = []interface{}{user.ID}

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

func (r *Repo) GetTask(taskID uint64) (models.Task, error) {
	var (
		task  dbTask
		query string
		err   error
	)

	query = `SELECT * FROM tasks WHERE id = ?`
	err = r.db.Get(&task, query, taskID)
	if err != nil {
		logger.LogError("Failed read task from db", "tasks/repo/mysql", strconv.FormatUint(taskID, 10), err)
		return models.Task{}, ErrTaskNotExist
	}

	return r.toModelTask(task), nil
}

func (r *Repo) InsertTaskTrack(tackTrack models.TaskTrack) error {
	var (
		dbTrack dbTaskTrack
		query   string
		err     error
	)

	dbTrack = r.toDbTaskTrack(tackTrack)

	if dbTrack.ID == 0 {
		query = `INSERT INTO task_track SET
		task_id = :task_id,
		user_id = :user_id,
		start_time = :start_time,
		end_time = :end_time,
		difference = :difference`
	} else {
		query = `UPDATE task_track SET
		task_id = :task_id,
		user_id = :user_id,
		start_time = :start_time,
		end_time = :end_time,
		difference = :difference
		WHERE id = :id`
	}

	_, err = r.db.NamedExec(query, &dbTrack)
	if err != nil {
		trackString := fmt.Sprintf("task_id: %d, user_id: %d, start_time: %s", dbTrack.TaskID, dbTrack.UserID, dbTrack.StartTime)
		logger.LogError(ErrFailedWriteTrack.Error(), "tasks/repo/mysql", trackString, err)
		return ErrFailedWriteTrack
	}

	return nil
}

func (r *Repo) GetLastTaskTrack(taskID, userID uint64) (models.TaskTrack, error) {
	var (
		dbTrack dbTaskTrack
		mTrack  models.TaskTrack
		query   string
		err     error
	)

	query = `SELECT * FROM task_track WHERE
	task_id = ? AND user_id = ?
	ORDER BY id DESC LIMIT 1`
	err = r.db.Get(&dbTrack, query, taskID, userID)
	if err != nil {
		logger.LogError(ErrFailedReadTrack.Error(), "tasks/repo/mysql", fmt.Sprintf("task_id: %d, user_id: %d", taskID, userID), err)
		return models.TaskTrack{}, ErrFailedReadTrack
	}

	mTrack = r.toModelTaskTrack(dbTrack)

	return mTrack, nil
}

func (r *Repo) Close() error {
	r.db.Close()
	return nil
}
