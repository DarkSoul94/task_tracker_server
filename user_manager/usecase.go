package user_manager

import (
	"github.com/DarkSoul94/task_tracker_server/models"
	"github.com/DarkSoul94/task_tracker_server/user_manager/perm_manager"
)

type UserManagerUC interface {
	CreateUser(user *models.User) (uint64, error)
	UserUpdate(author *models.User, userID, groupID uint64) error
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(id uint64) (models.User, error)
	GetUsersList(askUser *models.User) ([]models.User, error)

	GetGroupByID(user *models.User, groupID uint64) (models.Group, error)
	GetGroupList(user *models.User) ([]models.Group, error)
	GroupUpdate(id uint64, permission []byte) error
	CreateGroup(name string, permissions []byte) (uint64, error)

	GetFullPermListInBytes() (perm_manager.PermLayer, error)
}
