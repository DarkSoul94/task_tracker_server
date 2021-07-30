package mysql

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		db: sqlx.NewDb(db, "mysql"),
	}
}

func (r *Repo) Close() error {
	r.db.Close()
	return nil
}
