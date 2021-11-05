package usecase

import (
	"github.com/DarkSoul94/task_tracker_server/user_manager"
	"github.com/DarkSoul94/task_tracker_server/user_manager/permissions"
)

type Usecase struct {
	repo        user_manager.UserManagerRepo
	permManager permissions.PermManager
}
