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
	CreateDate  time.Time     `db:"create_time"`
	InWorkTime  time.Duration `db:"in_work_time"`
	AuthorID    uint64        `db:"author_id"`
	Developer   sql.NullInt64 `db:"developer"`
	StatusID    uint64        `db:"status_id"`
}
