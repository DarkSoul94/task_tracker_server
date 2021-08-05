package usecase

import (
	"github.com/DarkSoul94/task_tracker_server/models"
	"github.com/DarkSoul94/task_tracker_server/user_manager"
)

// NewUsecase ...
func NewUsecase(repo user_manager.UserManagerRepo) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) GetGroupByID(groupID uint64) (models.Group, error) {
	return u.repo.GetGroupByID(groupID)
}

func (u *Usecase) CreateUser(userName, passHash string) (models.User, error) {
	return u.repo.CreateUser(userName, passHash)
}

func (u *Usecase) GetUserByName(userName string) (models.User, error) {
	return u.repo.GetUserByName(userName)
}

func (u *Usecase) PermissionsCheck(group models.Group, targetActions ...string) map[string]interface{} {
	result := make(map[string]interface{})
	//TODO пересмотреть работу с доступами и целевыми действиями. Отойти от интерфейсов для сохранения строгой типизации
	for _, key := range targetActions {
		result[key] = permissionsChecksMap[key](group)
	}
	return result
}
