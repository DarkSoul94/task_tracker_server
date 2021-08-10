package user_manager

import "github.com/DarkSoul94/task_tracker_server/models"

type UserManagerRepo interface {
	CreateUser(userName, passHash string) (models.User, error)
	GetUserByName(name string) (models.User, error)
	GetUsersList() ([]models.User, error)

	GetGroupByID(groupID uint64) (models.Group, error)
	GetGroupList() ([]models.Group, error)

	Close() error
}
