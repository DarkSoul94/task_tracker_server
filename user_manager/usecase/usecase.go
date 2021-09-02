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

func (u *Usecase) CreateUser(user *models.User) (uint64, error) {
	return u.repo.CreateUser(user)
}

func (u *Usecase) GetUserByEmail(email string) (models.User, error) {
	return u.repo.GetUserByEmail(email)
}

func (u *Usecase) GetUserByID(id uint64) (models.User, error) {
	return u.repo.GetUserByID(id)
}

func (u *Usecase) GetUsersList(user *models.User) ([]models.User, error) {
	var (
		err       error
		userList  []models.User
		groupList []models.Group
	)
	if _, err = u.TargetActionPermissionCheck(user, user_manager.KeyGet); err != nil {
		return nil, err
	}

	if userList, err = u.repo.GetUsersList(); err != nil {
		return nil, ErrFailedGetUsersList
	}

	if groupList, err = u.repo.GetGroupList(); err != nil {
		return nil, ErrFailedGetGroupList
	}

	for _, user := range userList {
		for _, group := range groupList {
			if user.Group.ID == group.ID {
				user.Group = &group
				break
			}
		}
	}

	return userList, nil
}

func (u *Usecase) GetGroupByID(groupID uint64) (models.Group, error) {
	return u.repo.GetGroupByID(groupID)
}

func (u *Usecase) TargetActionPermissionCheck(user *models.User, actions ...string) (map[string]string, error) {
	var (
		group models.Group
		err   error
	)
	result := make(map[string]string)

	group, err = u.repo.GetGroupByID(user.Group.ID)
	if err != nil {
		return nil, err
	}
	permissions := group.ParsePermissionsByAction(actions...)
	for _, val := range actions {
		action := checksList[val](permissions)
		if len(action) == 0 {
			return nil, ErrUnauthorized
		}
		result[val] = action
	}
	return result, nil
}
