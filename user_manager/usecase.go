package user_manager

import "github.com/DarkSoul94/task_tracker_server/models"

type UserManagerUC interface {
	GetGroupByID(groupID uint64) (models.Group, error)
	CreateUser(userName, passHash string) (models.User, error)
	GetUserByName(userName string) (models.User, error)
	TargetActionPermissionCheck(group models.Group, actions ...string) map[string]string
}
