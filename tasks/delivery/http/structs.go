package http

import "github.com/DarkSoul94/task_tracker_server/tasks"

// Handler ...
type Handler struct {
	ucTasks tasks.TasksUC
}

type Responce struct {
	Status string      `json:"status"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

type newTask struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type outTask struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
