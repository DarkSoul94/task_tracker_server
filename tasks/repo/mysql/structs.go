package mysql

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

type dbTask struct {
	ID          uint64        `db:"id"`
	Name        string        `db:"name"`
	Description string        `db:"description"`
	CreateDate  time.Time     `db:"creation_date"`
	InWorkTime  time.Duration `db:"in_work_time"`
	AuthorID    uint64        `db:"author_id"`
	DeveloperID sql.NullInt64 `db:"developer_id"`
	CustomerID  sql.NullInt64 `db:"customer_id"`
	StatusID    uint64        `db:"status_id"`
	CategoryID  uint64        `db:"category_id"`
	ProjectID   sql.NullInt64 `db:"project_id"`
	Priority    bool          `db:"priority"`
	ExecOrder   uint64        `db:"exec_order"`
}

type dbTaskTrack struct {
	ID         uint64       `db:"id"`
	TaskID     uint64       `db:"task_id"`
	UserID     uint64       `db:"user_id"`
	StartTime  time.Time    `db:"start_time"`
	EndTime    sql.NullTime `db:"end_time"`
	Difference uint64       `db:"difference"`
}
