package mysql

import (
	"database/sql"
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

func (r *Repo) GetGroup(id uint64) (*models.Group, error) {
	var (
		group dbGroup
		query string
		err   error
	)

	query = `SELECT * FROM user_groups WHERE id = ?`
	err = r.db.Get(&group, query, id)
	if err != nil {
		logger.LogError(ErrReadGroup.Error(), "task_tracker_server/repo/mysql", strconv.FormatUint(id, 10), err)
		return nil, ErrReadGroup
	}

	return r.toModelGroup(group), nil
}

func (r *Repo) CreateUser(userName, passHash string) (models.User, error) {
	var (
		dbLoginUser dbLoginUser
		dbUser      dbUser
		query       string
		err         error
	)

	dbLoginUser = r.toDbLoginUser(userName, passHash)

	query = `INSERT INTO users SET
	name = :name,
	pass_hash = :pass_hash`

	_, err = r.db.NamedExec(query, dbLoginUser)
	if err != nil {
		logger.LogError(ErrCreateUser.Error(), "task_trecker_server/repo/mysql", "", err)
		return models.User{}, ErrCreateUser
	}

	query = `SELECT * FROM users WHERE name = ?`
	r.db.Get(&dbUser, query, userName)

	mUser := r.toModelUser(dbUser)
	mUser.Group, err = r.GetGroup(dbUser.GroupID)
	if err != nil {
		return models.User{}, err
	}

	return mUser, nil
}

func (r *Repo) GetUser(name string) (models.User, error) {
	var (
		dbUser dbUser
		query  string
		err    error
	)

	query = `SELECT * FROM users WHERE name = ?`
	err = r.db.Get(&dbUser, query, name)
	if err != nil {
		logger.LogError(ErrReadUser.Error(), "task_tracker_server/repo/mysql", name, err)
		return models.User{}, ErrReadUser
	}

	mUser := r.toModelUser(dbUser)
	mUser.Group, err = r.GetGroup(dbUser.GroupID)
	if err != nil {
		return models.User{}, err
	}

	return mUser, nil
}

func (r *Repo) Close() error {
	r.db.Close()
	return nil
}
