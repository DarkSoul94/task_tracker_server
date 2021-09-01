package http

import (
	"github.com/DarkSoul94/task_tracker_server/tasks"
	"github.com/gin-gonic/gin"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.RouterGroup, ucTasks tasks.TasksUC, middlewares ...gin.HandlerFunc) {
	h := NewHandler(ucTasks)

	tasksEndpoints := router.Group("/tasks")
	tasksEndpoints.Use(middlewares...)
	{
		//http://localhost:8585/task_tracker/tasks
		tasksEndpoints.POST("", h.CreateTask)
		//http://localhost:8585/task_tracker/tasks?user_id=1
		tasksEndpoints.GET("", h.GetTasksList)
		//http://localhost:8585/task_tracker/tasks/update
		tasksEndpoints.POST("/update", h.UpdateTask)
		//http://localhost:8585/task_tracker/tasks/track
		tasksEndpoints.POST("/track", h.TrackTask)
	}

	categoryEndpoints := router.Group("/category")
	categoryEndpoints.Use(middlewares...)
	{
		//http://localhost:8585/task_tracker/category/list
		categoryEndpoints.GET("/list", h.GetCategoryList)
	}

}
