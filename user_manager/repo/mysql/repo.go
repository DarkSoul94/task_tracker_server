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
	group, err := r.GetGroupByID(dbUser.GroupID)
	mUser.Group = &group
	if err != nil {
		return models.User{}, err
	}

	return mUser, nil
}

func (r *Repo) GetUserByName(name string) (models.User, error) {
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
	group, err := r.GetGroupByID(dbUser.GroupID)
	mUser.Group = &group
	if err != nil {
		return models.User{}, err
	}

	return mUser, nil
}

func (r *Repo) GetUsersList() ([]models.User, error) {
	var (
		dbUsersList []dbUser
		err         error
	)
	query := `SELECT * FROM users`

	if err = r.db.Select(&dbUsersList, query); err != nil {
		logger.LogError(ErrReadUsersList.Error(), "user_manager/repo/mysql", "", err)
		return []models.User{}, ErrReadUsersList
	}

	mUsersList := make([]models.User, 0)
	for _, val := range dbUsersList {
		mUsersList = append(mUsersList, r.toModelUser(val))
	}
	return mUsersList, nil
}

func (r *Repo) GetGroupByID(groupID uint64) (models.Group, error) {
	var (
		group dbGroup
		query string
		err   error
	)
	query = `SELECT * FROM user_groups WHERE id = ?`
	err = r.db.Get(&group, query, groupID)
	if err != nil {
		logger.LogError(ErrReadGroup.Error(), "user_manager/repo/mysql", strconv.FormatUint(groupID, 10), err)
		return models.Group{}, ErrReadGroup
	}

	return r.toModelGroup(group), nil
}

func (r *Repo) GetGroupList() ([]models.Group, error) {
	var (
		dbGroupsList []dbGroup
		err          error
	)
	query := `SELECT * FROM user_groups`

	if err = r.db.Select(&dbGroupsList, query); err != nil {
		logger.LogError(ErrReadGroupsList.Error(), "user_manager/repo/mysql", "", err)
		return []models.Group{}, ErrReadGroupsList
	}

	mGroupsList := make([]models.Group, 0)
	for _, val := range dbGroupsList {
		mGroupsList = append(mGroupsList, r.toModelGroup(val))
	}
	return mGroupsList, nil

}

func (r *Repo) Close() error {
	r.db.Close()
	return nil
}
