package user_manager

import "github.com/DarkSoul94/task_tracker_server/models"

type UserManagerUC interface {
	CreateUser(userName, passHash string) (models.User, error)
	GetUserByName(userName string) (models.User, error)
	GetUsersList(user *models.User) ([]models.User, error)

	GetGroupByID(groupID uint64) (models.Group, error)

	TargetActionPermissionCheck(user *models.User, actions ...string) (map[string]string, error)
}
