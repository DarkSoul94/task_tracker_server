package http

import "github.com/DarkSoul94/task_tracker_server/auth"

// Handler ...
type Handler struct {
	ucAuth auth.AuthUC
}

type Responce struct {
	Status string      `json:"status"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data"`
}

type inpGroup struct {
	ID   uint64 `json:"group_id,omitempty"`
	Name string `json:"group_name,omitempty"`
}

type loginUser struct {
	UserName string `json:"user_name"`
	Password []byte `json:"password"`
}

type inpUser struct {
	ID    uint64    `json:"user_id,omitempty"`
	Name  string    `json:"user_name,omitempty"`
	Token string    `json:"token"`
	Group *inpGroup `json:"group,omitempty"`
}
