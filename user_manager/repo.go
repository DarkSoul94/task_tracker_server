package user_manager

import "github.com/DarkSoul94/task_tracker_server/models"

type UserManagerRepo interface {
	GetGroupByID(groupID uint64) (models.Group, error)
	CreateUser(userName, passHash string) (models.User, error)
	GetUserByName(name string) (models.User, error)

	Close() error
}
