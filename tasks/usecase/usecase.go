package tasksUC

import (
	"time"

	"github.com/DarkSoul94/task_tracker_server/models"
	"github.com/DarkSoul94/task_tracker_server/tasks"
	"github.com/DarkSoul94/task_tracker_server/tasks/usecase/updTaskDisp"
	"github.com/DarkSoul94/task_tracker_server/user_manager"
)

// Usecase ...
type Usecase struct {
	repo        tasks.TasksRepo
	userManager user_manager.UserManagerUC
	updTaskDisp updTaskDisp.UpdTaskDisp
}

// NewUsecase ...
func NewUsecase(repo tasks.TasksRepo, userManager user_manager.UserManagerUC) *Usecase {
	return &Usecase{
		repo:        repo,
		userManager: userManager,
		updTaskDisp: updTaskDisp.NewTaskUpdDispatcher(repo, userManager),
	}
}

func (u *Usecase) CreateTask(task *models.Task) error {
	var err error
	//TODO add permissions check

	task.FillNewTask()
	err = u.repo.CreateTask(task)
	if err != nil {
		return ErrFailedCreateTask
	}
	return nil
}

func (u *Usecase) UpdateTask(user *models.User, task models.Task) error {
	var (
		err       error
		existTask *models.Task
	)
	existTask, err = u.repo.GetTask(task.ID)
	if err != nil {
		return err
	}

	u.updTaskDisp.SetData(existTask, &task, user)

	//TODO add permissions check

	return nil
}

func (u *Usecase) GetTasksList(user *models.User) ([]*models.Task, error) {
	var (
		taskList []*models.Task
		err      error
	)

	//TODO add permissions check

	if taskList, err = u.repo.GetTasksList(*user); err != nil {
		return []*models.Task{}, ErrFailedGetTasksList
	}

	return taskList, nil
}

func (u *Usecase) GetTask(taskID uint64) (*models.Task, error) {
	task, err := u.repo.GetTask(taskID)
	if err != nil {
		return &models.Task{}, err
	}

	err = u.fillTask(task)
	if err != nil {
		return &models.Task{}, err
	}

	return task, nil
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
		task.InWorkTime += time.Duration(track.Difference)
	}
	err = u.repo.InsertTaskTrack(track)
	if err != nil {
		return err
	}

	task.Tracked = status
	err = u.repo.UpdateTask(task)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) GetCategoryList() ([]*models.Category, error) {
	return u.repo.GetCategoryList()
}
