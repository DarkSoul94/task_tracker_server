package http

import (
	"github.com/DarkSoul94/task_tracker_server/tasks"
	"github.com/gin-gonic/gin"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.Engine, ucTasks tasks.TasksUC, middlewares ...gin.HandlerFunc) {
	h := NewHandler(ucTasks)

	tasksEndpoints := router.Group("/tasks")
	tasksEndpoints.Use(middlewares...)
	{
		//http://localhost:8585/tasks
		tasksEndpoints.POST("", h.CreateTask)
		//http://localhost:8585/tasks?user_id=1
		tasksEndpoints.GET("", h.GetTasksList)
		//http://localhost:8585/tasks/update
		tasksEndpoints.POST("/update", h.UpdateTask)
	}

}
