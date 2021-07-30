package http

import "github.com/DarkSoul94/task_tracker_server/auth"

// Handler ...
type Handler struct {
	ucAuth auth.AuthUC
}

type Responce struct {
	Status string      `json:"status"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

type inpGroup struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type loginUser struct {
	UserName string `json:"user_name"`
	Password []byte `json:"password"`
}

type inpUser struct {
	ID    uint64    `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	Group *inpGroup `json:"group,omitempty"`
}
