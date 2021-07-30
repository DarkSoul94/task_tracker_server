package usecase

import "github.com/DarkSoul94/task_tracker_server/user_manager"

type Usecase struct {
	repo user_manager.UserManagerRepo
}
