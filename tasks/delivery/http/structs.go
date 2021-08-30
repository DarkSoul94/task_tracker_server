package http

import "github.com/DarkSoul94/task_tracker_server/tasks"

// Handler ...
type Handler struct {
	ucTasks tasks.TasksUC
}

type Response struct {
	Status string      `json:"status"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

type newTask struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryID  uint64 `json:"category_id"`
	ProjectID   uint64 `json:"project_id,omitempty"`
	DeveloperID uint64 `json:"developer_id,omitempty"`
	Priority    bool   `json:"priority,omitempty"`
	CustomerID  uint64 `json:"customer_id,omitempty"`
	ExecOrder   uint64 `json:"exec_order,omitempty"`
}

type outTask struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
