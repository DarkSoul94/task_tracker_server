package updTaskDisp

import "github.com/DarkSoul94/task_tracker_server/models"

type UpdTaskDisp interface {
	SetData(existTask, newTask *models.Task, user *models.User)
}
