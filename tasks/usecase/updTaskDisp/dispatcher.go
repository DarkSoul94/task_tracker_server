package updTaskDisp

import (
	"github.com/DarkSoul94/task_tracker_server/models"
	"github.com/DarkSoul94/task_tracker_server/tasks"
	"github.com/DarkSoul94/task_tracker_server/user_manager"
)

type Dispatcher struct {
	repo      tasks.TasksRepo
	um        user_manager.UserManagerUC
	existTask *models.Task
	newTask   *models.Task
	user      *models.User
}

func NewTaskUpdDispatcher(repo tasks.TasksRepo, um user_manager.UserManagerUC) *Dispatcher {
	return &Dispatcher{
		repo: repo,
		um:   um,
	}
}

func (d *Dispatcher) SetData(existTask, newTask *models.Task, user *models.User) {
	d.existTask = existTask
	d.newTask = newTask
	d.user = user
}

func (d *Dispatcher) TaskUpdate() error {
	return nil
}
