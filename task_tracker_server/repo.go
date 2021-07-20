package task_tracker_server

import "github.com/DarkSoul94/task_tracker_server/models"

// Repository ...
type Repository interface {
	GetGroup(id uint64) (*models.Group, error)
	CreateUser(userName, passHash string) (models.User, error)
	GetUser(name string) (models.User, error)
	Close() error
}
