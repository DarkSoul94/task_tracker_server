package models

import (
	"github.com/DarkSoul94/task_tracker_server/user_manager/perm_manager"
)

type Group struct {
	ID          uint64
	Name        string
	Permissions perm_manager.PermLayer
}
