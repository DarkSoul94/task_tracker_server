package usecase

import (
	"github.com/DarkSoul94/task_tracker_server/models"
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

func (u *Usecase) SignUp(user *models.LoginUser) (models.User, error) {
	var (
		hash string
		err  error
	)
	hash, err = user.GetPassHash()
	if err != nil {
		return models.User{}, err
	}

	return u.repo.CreateUser(user.Name, hash)
}

func (u *Usecase) SignIn(inpUser *models.LoginUser) (models.User, error) {
	var (
		user models.User
		err  error
	)

	user, err = u.repo.GetUser(inpUser.Name)
	if err != nil {
		return models.User{}, err
	}

	if inpUser.VerifyPass(user.PassHash) {
		return user, nil
	} else {
		return models.User{}, ErrLoginFailed
	}
}
