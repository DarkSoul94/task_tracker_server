package task_tracker_server

import "github.com/DarkSoul94/task_tracker_server/models"

// Usecase ...
type Usecase interface {
	SignIn(inpUser *models.LoginUser) (models.User, error)
	SignUp(user *models.LoginUser) (models.User, error)
}
