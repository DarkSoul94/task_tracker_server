package auth

import (
	"context"

	"github.com/DarkSoul94/task_tracker_server/models"
)

// Usecase ...
type AuthUC interface {
	LDAPSignIn(email, password string) (models.User, string, error)
	GenerateToken(user *models.User) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*models.User, error)
}
