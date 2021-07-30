package http

import "github.com/DarkSoul94/task_tracker_server/task_tracker_server"

// Handler ...
type Handler struct {
	uc task_tracker_server.Usecase
}

type Responce struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
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
