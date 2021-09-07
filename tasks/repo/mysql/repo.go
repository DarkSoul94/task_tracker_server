package mysql

import (
	"database/sql"
	"fmt"
	"strconv"

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

func (r *Repo) GetTasksList(user models.User) ([]*models.Task, error) {
	var (
		query string
	)
	dbTasks := make([]dbTask, 0)
	mTasks := make([]*models.Task, 0)

	query = `SELECT 
				id, 
				name, 
				description,
				creation_date,
				in_work_time,
				status_id,
				priority,
				exec_order
			FROM tasks `

	err := r.db.Select(&dbTasks, query)
	if err != nil {
		logger.LogError("Failed get tasks list from db", "task/repo/mysql", "", err)
		return nil, err
	}
	if len(dbTasks) == 0 {
		return []*models.Task{}, nil
	}
	for _, task := range dbTasks {
		mTasks = append(mTasks, r.toModelTask(task))
	}
	return mTasks, nil
}

func (r *Repo) GetTask(taskID uint64) (*models.Task, error) {
	var (
		task  dbTask
		query string
		err   error
	)

	query = `SELECT * FROM tasks WHERE id = ?`
	err = r.db.Get(&task, query, taskID)
	if err != nil {
		logger.LogError("Failed read task from db", "tasks/repo/mysql", strconv.FormatUint(taskID, 10), err)
		return nil, ErrTaskNotExist
	}

	return r.toModelTask(task), nil
}

func (r *Repo) InsertTaskTrack(tackTrack *models.TaskTrack) error {
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
		logger.LogError(ErrAddTrack.Error(), "tasks/repo/mysql", trackString, err)
		return ErrAddTrack
	}

	return nil
}

func (r *Repo) GetLastTaskTrack(taskID, userID uint64) (*models.TaskTrack, error) {
	var (
		dbTrack dbTaskTrack
		mTrack  *models.TaskTrack
		query   string
		err     error
	)

	query = `SELECT * FROM task_track WHERE
	task_id = ? AND user_id = ?
	ORDER BY id DESC LIMIT 1`
	err = r.db.Get(&dbTrack, query, taskID, userID)
	if err != nil {
		logger.LogError(ErrGetTrack.Error(), "tasks/repo/mysql", fmt.Sprintf("task_id: %d, user_id: %d", taskID, userID), err)
		return nil, ErrGetTrack
	}

	mTrack = r.toModelTaskTrack(dbTrack)

	return mTrack, nil
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

func (r *Repo) GetCategoryList() ([]*models.Category, error) {
	var (
		categories  []dbCategory
		mCategories []*models.Category
		query       string
		err         error
	)

	query = `SELECT * FROM categories`
	err = r.db.Select(&categories, query)
	if err != nil {
		logger.LogError(ErrGetCategoryList.Error(), "tasks/repo/mysql", "", err)
		return nil, ErrGetCategoryList
	}

	for _, category := range categories {
		mCategories = append(mCategories, r.toModelCategory(category))
	}

	return mCategories, nil
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
