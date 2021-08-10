package http

import (
	"github.com/DarkSoul94/task_tracker_server/user_manager"
	"github.com/gin-gonic/gin"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.RouterGroup, uc user_manager.UserManagerUC, middlewares ...gin.HandlerFunc) {
	h := NewHandler(uc)

	userEndpoints := router.Group("/user")
	userEndpoints.Use(middlewares...)
	{
		//http://localhost:8585/task_tracker/user/update
		userEndpoints.POST("/update", h.UpdateUser)
		//http://localhost:8585/task_tracker/user/list
		userEndpoints.GET("/list", h.GetUsersList)
	}

	groupEndpoints := router.Group("/group")
	groupEndpoints.Use(middlewares...)
	{
		//http://localhost:8585/task_tracker/group/create
		groupEndpoints.POST("/create", h.CreateGroup)
		//http://localhost:8585/task_tracker/group/update
		groupEndpoints.POST("/update", h.UpdateGroup)
		//http://localhost:8585/task_tracker/group/list
		groupEndpoints.GET("/list", h.GetGroupsList)
	}

}
