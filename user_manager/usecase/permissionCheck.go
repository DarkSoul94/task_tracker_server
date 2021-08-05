package usecase

import (
	"github.com/DarkSoul94/task_tracker_server/models"
	"github.com/DarkSoul94/task_tracker_server/tasks"
)

var (
	permissionsChecksMap map[string]func(models.Group) interface{} = map[string]func(models.Group) interface{}{
		tasks.KeyGetTasks: checkForGetTaskList,
	}
)

func checkForGetTaskList(group models.Group) interface{} {
	return tasks.TargetAction_GetAllTasks
}
