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

func (u *Usecase) GetTasksList(user *models.User) ([]*models.Task, error) {
	var (
		taskList []*models.Task
		actions  map[string]string
		err      error
	)

	if actions, err = u.userManager.TargetActionPermissionCheck(user, tasks.KeyGet); err != nil {
		return []*models.Task{}, err
	}

	if taskList, err = u.repo.GetTasksList(actions[tasks.KeyGet], *user); err != nil {
		return []*models.Task{}, ErrFailedGetTasksList
	}

	for _, val := range taskList {
		err = u.fillTask(val)
		if err != nil {
			return []*models.Task{}, err
		}
	}
	return taskList, nil
}

func (u *Usecase) TrackTask(taskId uint64, user *models.User, status bool) error {
	var (
		task  *models.Task
		track *models.TaskTrack
		err   error
	)

	task, err = u.repo.GetTask(taskId)
	if err != nil {
		return err
	}

	if task.Developer == nil {
		return ErrNotDeveloper
	}

	if task.Developer.ID != user.ID {
		return ErrNotDeveloper
	}

	/*TODO: раскоментить когда будет апдейт таска с возможностью менять статус
	if task.Status.ID != models.TaskStatusMap[models.KeyTSInWork].ID {
		return ErrNotInWork
	}*/

	track, err = u.repo.GetLastTaskTrack(taskId, user.ID)
	if status {
		if err == nil {
			if track.EndTime.IsZero() {
				return ErrIsTracking
			}
		}

		track.StartTrack(taskId, user)
	} else {
		if err != nil || !track.EndTime.IsZero() {
			return ErrIsNotTracking
		}
		track.EndTrack()
		//TODO: add track.difference to task.inworktime and update task
	}
	err = u.repo.InsertTaskTrack(track)
	if err != nil {
		return err
	}

	return nil
}
func (u *Usecase) fillTask(task *models.Task) error {
	var (
		cat *models.Category
		pr  *models.Project
		err error
	)
	cat, err = u.repo.GetCategoryByID(task.Category.ID)
	if err != nil {
		return err
	}
	task.Category = cat

	if task.Author != nil {
		user, err := u.userManager.GetUserByID(task.Author.ID)
		if err != nil {
			return err
		}
		task.Author = &user
	}

	if task.Developer != nil {
		user, err := u.userManager.GetUserByID(task.Developer.ID)
		if err != nil {
			return err
		}
		task.Developer = &user
	}

	if task.Customer != nil {
		user, err := u.userManager.GetUserByID(task.Customer.ID)
		if err != nil {
			return err
		}
		task.Customer = &user
	}

	if task.Project != nil {
		pr, err = u.repo.GetProjectByID(task.Project.ID)
		if err != nil {
			return err
		}
	}
	task.Project = pr

	return nil
}

func (u *Usecase) GetCategoryList() ([]*models.Category, error) {
	return u.repo.GetCategoryList()
}
