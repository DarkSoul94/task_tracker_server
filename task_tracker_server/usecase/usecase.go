package usecase

import (
	"context"

	"github.com/DarkSoul94/task_tracker_server/task_tracker_server"
)

// Usecase ...
type Usecase struct {
	repo task_tracker_server.Repository
}

// NewUsecase ...
func NewUsecase(repo task_tracker_server.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

// HelloWorld ...
func (u *Usecase) HelloWorld(c context.Context) {
	println("Hello")
}
