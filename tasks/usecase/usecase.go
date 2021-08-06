package tasksUC

import (
	"github.com/DarkSoul94/task_tracker_server/models"
	"github.com/DarkSoul94/task_tracker_server/tasks"
	"github.com/DarkSoul94/task_tracker_server/user_manager"
)

// Usecase ...
type Usecase struct {
	repo        tasks.TasksRepo
	userManager user_manager.UserManagerUC
}

// NewUsecase ...
func NewUsecase(repo tasks.TasksRepo, userManager user_manager.UserManagerUC) *Usecase {
	return &Usecase{
		repo:        repo,
		userManager: userManager,
	}
}

func (u *Usecase) CreateTask(task models.Task) error {
	var err error

	task.FillNewTask()

	err = u.repo.CreateTask(task)
	if err != nil {
		return ErrFailedCreateTask
	}
	return nil
}

func (u *Usecase) GetTasksList(user *models.User) ([]models.Task, error) {
	var (
		err      error
		taskList []models.Task
	)

	actions := make(map[string]string)

	//TODO добавить поверку доступов
	taskList, err = u.repo.GetTasksList(actions[tasks.KeyGet], *user)
	if err != nil {
		return []models.Task{}, ErrFailedGetTasksList
	}

	return taskList, nil
}
