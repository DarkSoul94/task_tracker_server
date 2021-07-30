package auth

import "github.com/DarkSoul94/task_tracker_server/models"

// Usecase ...
type AuthUC interface {
	SignIn(inpUser *models.LoginUser) (string, models.User, error)
	SignUp(user *models.LoginUser) (models.User, error)
}
