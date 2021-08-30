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

func (r *Repo) GetTasksList(key string, user models.User) ([]*models.Task, error) {
	var (
		query string
	)

	args := make([]interface{}, 0)
	dbTasks := make([]dbTask, 0)
	mTasks := make([]*models.Task, 0)

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

func (r *Repo) GetCategoryByID(id uint64) (*models.Category, error) {
	var (
		mCat  *models.Category
		dbCat dbCategory
		err   error
	)
	err = r.db.Get(&dbCat, `
		SELECT * FROM categories
		WHERE id = ?
		`, id)
	if err != nil {
		logger.LogError("Failed get category from db", "task/repo/mysql", fmt.Sprintf("category id: %d", id), err)
		return nil, err
	}
	mCat = r.toModelCategory(dbCat)
	return mCat, nil
}

func (r *Repo) GetProjectByID(id uint64) (*models.Project, error) {
	var (
		mProject  *models.Project
		dbProject dbProject
		err       error
	)
	err = r.db.Get(&dbProject, `
		SELECT * FROM projects
		WHERE id = ?
		`, id)
	if err != nil {
		logger.LogError("Failed get category from db", "task/repo/mysql", fmt.Sprintf("project id: %d", id), err)
		return nil, err
	}
	mProject = r.toModelProject(dbProject)
	return mProject, nil
}

func (r *Repo) Close() error {
	r.db.Close()
	return nil
}
