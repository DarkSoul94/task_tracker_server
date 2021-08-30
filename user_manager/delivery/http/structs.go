package http

import (
	"github.com/DarkSoul94/task_tracker_server/user_manager"
)

// Handler ...
type Handler struct {
	ucUserManager user_manager.UserManagerUC
}

type Response struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

type outUser struct {
	ID    uint64    `json:"id"`
	Name  string    `json:"name"`
	Group *outGroup `json:"group"`
}

type outGroup struct {
	ID          uint64              `json:"id"`
	Name        string              `json:"name"`
	Permissions map[string][]string `json:"permissions"`
}
