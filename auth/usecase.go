package auth

import (
	"context"

	"github.com/DarkSoul94/task_tracker_server/models"
)

// Usecase ...
type AuthUC interface {
	SignIn(inpUser *models.LoginUser) (models.User, error)
	SignUp(user *models.LoginUser) (models.User, error)
	GenerateToken(user *models.User) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*models.User, error)
}
