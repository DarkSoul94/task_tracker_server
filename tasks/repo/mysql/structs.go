package mysql

import "github.com/jmoiron/sqlx"

type Repo struct {
	db *sqlx.DB
}

type dbGroup struct {
	ID   uint64 `db:"id"`
	Name string `db:"group_name"`
}

type dbLoginUser struct {
	Name     string `db:"name"`
	PassHash string `db:"pass_hash"`
}

type dbUser struct {
	ID       uint64 `db:"id"`
	Name     string `db:"name"`
	PassHash string `db:"pass_hash"`
	GroupID  uint64 `db:"group_id"`
}
