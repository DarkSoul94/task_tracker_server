package user_manager

import "github.com/DarkSoul94/task_tracker_server/models"

type UserManagerUC interface {
	CreateUser(user *models.User) (uint64, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(id uint64) (models.User, error)
	GetUsersList(askUser *models.User) ([]models.User, error)

	GetGroupByID(groupID uint64) (models.Group, error)

	TargetActionPermissionCheck(user *models.User, actions ...string) (map[string]string, error)
}
