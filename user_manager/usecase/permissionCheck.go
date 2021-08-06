package usecase

import (
	"github.com/DarkSoul94/task_tracker_server/models"
	"github.com/DarkSoul94/task_tracker_server/tasks"
)

var (
	checksList map[string]func(models.Group) string = map[string]func(models.Group) string{
		tasks.KeyGet: getTasks,
	}
)

func (u *Usecase) TargetActionPermissionCheck(user models.User, actions ...string) (map[string]string, error) {
	var (
		group models.Group
		err   error
	)
	result := make(map[string]string)

	group, err = u.repo.GetGroupByID(user.Group.ID)
	if err != nil {
		return nil, err
	}

	for _, val := range actions {
		result[val] = checksList[val](group)
	}
	return result, nil
}

func getTasks(group models.Group) string {

	return ""
}
